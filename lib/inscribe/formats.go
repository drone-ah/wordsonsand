package inscribe

import "gopkg.in/yaml.v2"

type UnmarshalFunc func(in []byte, out any) error
type MarshalFunc func(in any) ([]byte, error)

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
