package main

import (
	"bytes"
	"os"

	"github.com/adrg/frontmatter"
)

type Video struct {
	path    string
	meta    Metadata
	content []byte
}

func NewVideo(path string) (Video, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Video{}, nil
	}

	var meta Metadata
	content, err := frontmatter.Parse(bytes.NewReader(data), &meta)
	if err != nil {
		return Video{}, nil
	}

	return Video{
		path:    path,
		meta:    meta,
		content: content,
	}, nil
}
