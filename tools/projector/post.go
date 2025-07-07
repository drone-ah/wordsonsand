package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/drone-ah/wordsonsand/lib/inscribe"
)

type Video struct {
	path         string
	sourceRoot   string
	renderedPath string
	meta         Metadata
	scribed      inscribe.Scribed
}

func NewVideo(sourcePath string, sourceRoot string) (Video, error) {
	scribed, err := inscribe.NewScribedFromFile(sourcePath)
	if err != nil {
		return Video{}, nil
	}

	var meta Metadata = Metadata{
		Hashes: make(map[string]string),
	}

	err = scribed.FrontMatter(&meta)
	if err != nil {
		return Video{}, nil
	}

	return Video{
		path:       sourcePath,
		sourceRoot: sourceRoot,
		meta:       meta,
		scribed:    scribed,
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

func (v *Video) save() error {
	file, err := os.Create(v.path)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			slog.Warn("Unable to close file", "file", v.path, "error", err)
		}
	}()

	err = v.scribed.Write(v.meta, file)
	if err != nil {
		return err
	}

	return nil
}
