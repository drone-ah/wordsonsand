---
title: "Inscribe: Updating frontmatter in-place with Go and yaml.Node"
date: 2025-07-04T16:44:02+01:00
categories:
  - golang
tags:
  - golang
  - inscribe
  - yaml
  - frontmatter
  - tooling
  - markdown
  - projector
---

While building
[a little tool to sync up YouTube video descriptions from hugo](../wordsonsand/projector-sync.md),
I needed a library to read and write frontmatter in yaml.

I started with [adrg/frontmatter](https://github.com/adrg/frontmatter) before I
realised that it didn't have the ability to write back.

I considered contributing to that one, but writing back is a little more complex
than reading - particularly because the `frontmatter.Parse` in that one is built
to support partial reading.

Because adrg/frontmatter only unmarshals into a struct, and doesn’t store the
original bytes, you can’t write back without losing untouched keys.

Looking around, I could not find another frontmatter library for Go. I know that
python has a decent library which supports writing back to it (I used it in the
[depatcher](../wordsonsand/despatches.md]) but I don't want to write python.

<!-- more -->

## Multiple Formats

Let's make it extendable by defining a `Format` that will allow us to add in
other formats later:

```go
// A Format knows how to (un)marshal a particular format of frontmatter.
// e.g. yaml, toml etc.
type Format struct {
	// TODO: We could add details like delimiter to support other formats
	// and auto detection of format
	Unmarshal UnmarshalFunc
	Marshal   MarshalFunc
}

var yamlFormat = Format{
	Unmarshal: yaml.Unmarshal,
	Marshal:   yaml.Marshal,
}
```

## Writing Back

We could also do with a struct to hold the whole file contents so that we can
write it back easier.

```go
// A Scribed is a representation of a file that contains frontmatter and markdown content
type Scribed struct {
	format      Format
	frontmatter []byte
	Content     string
}
```

By storing the full frontmatter, we can later accept partial updates without
losing other keys.

## Naive Merging of Updates

We (the user) should be able to update just the keys we care about. All the
other keys should be preserved.

The easiest way I could find to do this was to Marshal, Unmarshall and then
merge with the raw Unmarshall:

### Minimal merge strategy (loses order and formatting)

```go
// Merge frontmatter
var raw map[string]any // full unmarshalled frontmatter
err := s.format.Unmarshal(s.frontmatter, &raw)

updatedBytes, _ := yaml.Marshal(fm) // convert updated to yaml

var updates map[string]any
yaml.Unmarshal(updatedBytes, &updates) // get updated keys as map

for k, v := range updates {
    raw[k] = v // overwrite only touched fields
}

// raw is now the preserved + updated keys
data, err := s.format.Marshal(raw)
if err != nil {
    return err
}
```

> ⚠️ **Warning**: Key ordering is lost
>
> Due to the way maps work, the key ordering is lost More accurately, the keys
> are sorted alphabetically during write.
>
> It fully rewrites the frontmatter, which also means that double quotes might
> disappear etc.

## In Place Merging of Updates

If it is important to keep the frontmatter formatting as much as possible, we
need to bigger sledgehammer.

I explored an approach using yaml.Node with ChatGPT's help.

I fitted it into the `Format` as well

```go
type MergeFunc func(raw []byte, fm any) ([]byte, error)

type Format struct {
	// TODO: We could add details like delimiter to support other formats
	// and auto detection of format
	Unmarshal UnmarshalFunc
	Marshal   MarshalFunc
	Merge     MergeFunc
}

func MergeYaml(raw []byte, fm any) ([]byte, error) {
	var node yaml.Node
	if err := yaml.Unmarshal(raw, &node); err != nil {
		return nil, err
	}

	// Expecting a document with a single mapping node
	if len(node.Content) == 0 || node.Content[0].Kind != yaml.MappingNode {
		return nil, errors.New("invalid frontmatter")
	}

	b, err := yaml.Marshal(fm)
	if err != nil {
		return nil, err
	}

	var updates map[string]string
	yaml.Unmarshal(b, &updates)

	m := node.Content[0]
	for key, value := range updates {
		// Search for the key and update it, or append a new one
		found := false
		for i := 0; i < len(m.Content); i += 2 {
			k := m.Content[i]
			if k.Value == key {
				m.Content[i+1].Value = value
				found = true
				break
			}
		}
		if !found {
			m.Content = append(m.Content,
				&yaml.Node{Kind: yaml.ScalarNode, Value: key},
				&yaml.Node{Kind: yaml.ScalarNode, Value: value},
			)
		}
	}

	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	if err := enc.Encode(&node); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
```

This preserves frontmatter formatting well - but it only handles flat YAML.
Nested maps require a bit more work.

## Supporting maps etc. as values

We need to switch to `map[string]any` and tweak a bit of the loop

```go
var updates map[string]any
yaml.Unmarshal(b, &updates)

m := node.Content[0]
for key, value := range updates {
    // Encode value to a yaml.Node
    valNode := &yaml.Node{}
    if err := valNode.Encode(value); err != nil {
        return nil, err
    }

    // Search and replace or append
    found := false
    for i := 0; i < len(m.Content); i += 2 {
        if m.Content[i].Value == key {
            m.Content[i+1] = valNode
            found = true
            break
        }
    }
    if !found {
        m.Content = append(m.Content,
            &yaml.Node{Kind: yaml.ScalarNode, Value: key},
            valNode,
        )
    }
}

```

It's funny how things can be more complicated than it first seems.

## Minor tweaks

I also wanted to add the delimiter into the format and use that instead of
hardcoding, which was easy enough.

```go
// A Format knows how to (un)marshal a particular format of frontmatter.
// e.g. yaml, toml etc.
type Format struct {
	Delimiter string
	Unmarshal UnmarshalFunc
	Marshal   MarshalFunc
	Merge     MergeFunc
}

var yamlFormat = Format{
	Delimiter: "---",
	Unmarshal: yaml.Unmarshal,
	Marshal:   yaml.Marshal,
	Merge:     MergeYaml,
}

func (s *Scribed) Write(fm any, out io.Writer) error {
	// Merge frontmatter
	data, err := s.format.Merge(s.frontmatter, fm)
	if err != nil {
		return err
	}

	io.WriteString(out, s.format.Delimiter+"\n")
	out.Write(data)
	io.WriteString(out, s.format.Delimiter+"\n\n")
	io.WriteString(out, s.Content)

	return nil
}

// splitFrontmatter will split frontmatter from Content and store them
func (s *Scribed) splitFrontmatter(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	parts := bytes.SplitN(data, []byte("\n"+s.format.Delimiter+"\n"), 2)
	if len(parts) != 2 {
		return errors.New("invalid frontmatter format")
	}

	// Remove the opening '---\n' from the first part
	s.frontmatter = bytes.TrimPrefix(parts[0], []byte("---\n"))

	s.Content = string(bytes.TrimLeft(parts[1], "\r\n"))

	return nil
}
```

## Conclusion

This is probably the main bits of functionality I'll need to continue with the
[projector sync](../wordsonsand/projector-sync.md).

I had considered `frontmatter` to be a blackbox with complicated functionality,
but cutting it up and working on it has demystified it and made it easier to
work with. ChatGPT helped.

It should be fairly straightforward to add other frontmatter formats like TOML,
and to autodetect the formats, but I don't need it right now.

## Links

- [Source Code on GitHub](https://github.com/drone-ah/wordsonsand/tree/main/lib/inscribe)
