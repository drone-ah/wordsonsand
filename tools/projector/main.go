package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli/v3"
)

type Metadata struct {
	Title       string            `yaml:"title"`
	VideoId     string            `yaml:"youtubeId"`
	PublishDate string            `yaml:"publishDate"`
	Hashes      map[string]string `yaml:"hashes"`
}

func main() {
	// Use regular logging (through slog, so we get the levels)
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
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
			},
			{
				Name:  "sync",
				Usage: "sync metadata",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "source",
						Usage: "location of yaml info",
					},
					&cli.StringFlag{
						Name:  "rendered",
						Usage: "location of rendered descriptions",
					},
					&cli.StringFlag{
						Name:    "client-id",
						Usage:   "google oauth client id",
						Sources: cli.EnvVars("GOOGLE_CLIENT_ID"),
					},
					&cli.StringFlag{
						Name:    "client-secret",
						Usage:   "google oauth client secret",
						Sources: cli.EnvVars("GOOGLE_CLIENT_SECRET"),
					},
					&cli.StringFlag{
						Name:    "refresh-token",
						Usage:   "google refresh token",
						Sources: cli.EnvVars("GOOGLE_REFRESH_TOKEN"),
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return sync(cmd)
				},
			},
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

	targetRenderedDir, err := getTargetDir(renderedPath)
	if err != nil {
		return nil
	}

	videos, err := findRecentVideos(targetSourceDir)
	for _, video := range videos {
		_, err := video.getDescription(targetRenderedDir)
		if err != nil {
			slog.Warn("unable to find rendered file", "file", video.renderedPath)
		}
	}
	return nil
}

func sync(cmd *cli.Command) error {
	sourcePath := cmd.String("source")
	renderedPath := cmd.String("rendered")
	clientId := cmd.String("client-id")
	clientSecret := cmd.String("client-secret")
	refreshToken := cmd.String("refresh-token")

	targetSourceDir, err := getTargetDir(sourcePath)
	if err != nil {
		return err
	}

	targetRenderedDir, err := getTargetDir(renderedPath)
	if err != nil {
		return nil
	}

	vService, err := NewYouTube(clientId, clientSecret, refreshToken)
	if err != nil {
		return err
	}

	videos, err := findRecentVideos(targetSourceDir)
	for _, video := range videos {
		bdesc, err := video.getDescription(targetRenderedDir)
		if err != nil {
			slog.Warn("unable to find rendered file", "file", video.renderedPath)
		}

		// We want to hash the contents of description
		// Check with the hash in the metadata to see if it matches
		hash := md5.Sum(bdesc)
		strHash := hex.EncodeToString(hash[:])

		if video.meta.Hashes["description"] != strHash {
			// update description
			// call youtube api to update description

			vId := video.meta.VideoId
			if vId != "" {
				err := vService.UpdateDescription(video.meta.VideoId, string(bdesc))
				if err != nil {
					slog.Warn("Unable to set description", "err", err)
				}
			}

			slog.Debug("updated description", "file", video.path)
			// update the hash in the source file
			video.meta.Hashes["description"] = strHash

			err := video.save()
			if err != nil {
				slog.Warn("unable to save file", "file", video.path)
			}
		}

	}
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

func findRecentVideos(sourceRoot string) ([]Video, error) {
	now := time.Now().UTC()
	cutoff := now.AddDate(0, 0, -30)

	var posts []Video

	err := filepath.Walk(sourceRoot, func(fullPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		slog.Debug("checking:", "path", fullPath)

		post, err := NewVideo(fullPath, sourceRoot)
		if err != nil {
			return err
		}

		meta := post.meta
		if meta.PublishDate != "" {
			pd, err := time.Parse(time.RFC3339, meta.PublishDate)
			if err != nil {
				slog.Warn("unable to parse", "path", fullPath, "publishDate", meta.PublishDate)
			}
			if err == nil && pd.After(cutoff) {
				posts = append(posts, post)
				slog.Debug("filtered", "path", fullPath)
			}
		}
		return nil
	})

	return posts, err
}
