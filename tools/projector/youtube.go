package main

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YouTube struct{}

func (y YouTube) UpdateDescription(videoId string, desc string) error {
	ctx := context.Background()
	ytService, err := youtube.NewService(ctx, option.WithAPIKey("AIzaSyDolWdM6EDARd1dVw2TU5Ef4GsqYV5Co2g"))
	if err != nil {
		return err
	}

	ytVideo := youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Description: desc,
		},
	}

	call := ytService.Videos.Update([]string{"snippet.description"}, &ytVideo)
	_, err = call.Do()

	return err
}
