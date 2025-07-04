package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli/v3"
)

type Metadata struct {
	Title       string `yaml:"title"`
	PublishDate string `yaml:"publishDate"`
}

func main() {
	// Use regular logging (through slog, so we get the levels)
	handler := slog.NewTextHandler(os.Stderr, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	cmd := &cli.Command{
		Name:  "validate",
		Usage: "validate metadata",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "source",
				Usage: "location of yaml info",
			},
			&cli.StringFlag{
				Name:  "rendered",
				Usage: "location of rendered descriptions",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return validate(cmd.String("source"), cmd.String("rendered"))
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func validate(sourcePath string, renderedPath string) error {
	targetSourceDir, err := getTargetDir(sourcePath)
	if err != nil {
		return err
	}

	_, err = findRecentVideos(targetSourceDir)
	return nil
}

func getTargetDir(basePath string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	baseDir := os.Getenv("BUILD_WORKING_DIRECTORY")
	if baseDir == "" {
		baseDir = cwd
	}
	targetDir := filepath.Join(baseDir, basePath)
	return targetDir, nil
}

func findRecentVideos(path string) ([]Video, error) {
	now := time.Now().UTC()
	cutoff := now.AddDate(0, 0, -30)

	var posts []Video

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		slog.Debug("checking:", "path", path)

		post, err := NewVideo(path)
		if err != nil {
			return err
		}

		meta := post.meta
		if meta.PublishDate != "" {
			pd, err := time.Parse(time.RFC3339, meta.PublishDate)
			if err != nil {
				slog.Warn("unable to parse", "path", path, "publishDate", meta.PublishDate)
			}
			if err == nil && pd.After(cutoff) {
				posts = append(posts, post)
				slog.Debug("filtered", "path", path)
			}
		}
		return nil
	})

	return posts, err
}
