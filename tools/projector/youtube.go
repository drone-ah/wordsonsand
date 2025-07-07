package main

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YouTube struct {
	ClientId     string
	ClientSecret string
	RefreshToken string
}

func (y YouTube) UpdateDescription(videoId string, desc string) error {

	conf := &oauth2.Config{
		ClientID:     y.ClientId,
		ClientSecret: y.ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/youtube"},
	}

	// Construct a token from just the refresh token
	token := &oauth2.Token{RefreshToken: y.RefreshToken}

	ctx := context.Background()

	// Create an authenticated client
	httpClient := conf.Client(ctx, token)

	ytService, err := youtube.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return err
	}

	ytVideo := youtube.Video{
		Id: videoId,
		Snippet: &youtube.VideoSnippet{
			Title:       "some title",
			CategoryId:  "20",
			Description: desc,
		},
	}

	call := ytService.Videos.Update([]string{"snippet"}, &ytVideo)
	_, err = call.Do()

	return err
}
