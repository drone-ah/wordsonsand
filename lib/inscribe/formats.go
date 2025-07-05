package inscribe

import (
	"bytes"
	"errors"

	"gopkg.in/yaml.v3"
)

type UnmarshalFunc func(in []byte, out any) error
type MarshalFunc func(in any) ([]byte, error)
type MergeFunc func(raw []byte, fm any) ([]byte, error)

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

	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(&node); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
