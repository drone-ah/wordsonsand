package main

import (
	"context"
	"fmt"
	"log"
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

func validate(pathParam string, renderedPath string) error {
	targetDir, err := getTargetDir(pathParam)
	if err != nil {
		return err
	}

	_, err = findFiles(targetDir)
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

func findFiles(path string) ([]Post, error) {
	now := time.Now().UTC()
	cutoff := now.AddDate(0, 0, -30)

	var posts []Post

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)

		post, err := NewPost(path)
		if err != nil {
			return err
		}

		meta := post.meta
		if meta.PublishDate != "" {
			pd, err := time.Parse(time.RFC3339, meta.PublishDate)
			if err == nil && pd.After(cutoff) {
				posts = append(posts, post)
				fmt.Printf("filtered: %s\n", path)
			}
		}
		return nil
	})

	return posts, err
}
