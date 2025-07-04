package main

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/adrg/frontmatter"
)

type Video struct {
	path         string
	sourceRoot   string
	renderedPath string
	meta         Metadata
	content      []byte
}

func NewVideo(sourcePath string, sourceRoot string) (Video, error) {
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return Video{}, nil
	}

	var meta Metadata
	content, err := frontmatter.Parse(bytes.NewReader(data), &meta)
	if err != nil {
		return Video{}, nil
	}

	return Video{
		path:       sourcePath,
		sourceRoot: sourceRoot,
		meta:       meta,
		content:    content,
	}, nil
}

func (v *Video) getDescription(renderedRoot string) ([]byte, error) {
	relPath := v.path[len(v.sourceRoot):]
	relPath = relPath[:len(relPath)-3] // trim .md at the end
	v.renderedPath = filepath.Join(renderedRoot, relPath, "index.txt")
	slog.Debug("paths", "relative", relPath, "rendered", v.renderedPath)
	b, err := os.ReadFile(v.renderedPath)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}
