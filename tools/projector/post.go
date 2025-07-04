package main

import (
	"bytes"
	"os"

	"github.com/adrg/frontmatter"
)

type Post struct {
	meta    Metadata
	content []byte
}

func NewPost(path string) (Post, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Post{}, nil
	}

	var meta Metadata
	content, err := frontmatter.Parse(bytes.NewReader(data), &meta)
	if err != nil {
		return Post{}, nil
	}

	return Post{
		meta:    meta,
		content: content,
	}, nil
}
