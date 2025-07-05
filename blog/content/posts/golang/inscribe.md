---
title: "Inscribe: A simple golang frontmatter parser"
date: 2025-07-04T16:44:02+01:00
draft: true
---

While building
[a little tool to sync up YouTube video descriptions from hugo](../wordsonsand/projector-sync.md),
I needed a library to read and write frontmatter in yaml.

I started with [adrg/frontmatter](https://github.com/adrg/frontmatter) before I
realised that it didn't have the ability to write back.

I considered contributing to that one, but writing back is a little more complex
than reading - particularly because the `frontmatter.Parse` in that one is built
to support partial reading.

A `Write` wouldn't be super hard to write, but the problem is that the full
frontmatter isn't stored anywhere. As such, if you write back a partially read
frontmatter, you'd lose the other keys.

Looking around, I could not find another frontmatter library for golang. I know
that python has a decent library which supports writing back to it (I used it in
the [depatcher](../wordsonsand/despatches.md]) but I don't want to write python.

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

By storing the full frontmatter, when we get a partial one back, we can merge
it, and write the whole thing back.

## Merging Updates

We want the user (in this case also us), to be able to update only the keys they
are interested in. All the other keys should be preserved.

The easiest way I could find to do this was to Marshal, Unmarshall and then
merge with the raw Unmarshall:

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
> end up ordered. I could not figure out an easy way to preserve ordering.
