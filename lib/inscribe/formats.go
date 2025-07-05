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
	// TODO: We could add details like delimiter to support other formats
	// and auto detection of format
	Unmarshal UnmarshalFunc
	Marshal   MarshalFunc
	Merge     MergeFunc
}

var yamlFormat = Format{
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
