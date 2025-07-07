package main

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YouTube struct {
	service *youtube.Service
}

func NewYouTube(ClientId string, ClientSecret string, RefreshToken string) (YouTube, error) {

	conf := &oauth2.Config{
		ClientID:     ClientId,
		ClientSecret: ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/youtube"},
	}

	// Construct a token from just the refresh token
	token := &oauth2.Token{RefreshToken: RefreshToken}

	ctx := context.Background()

	// Create an authenticated client
	httpClient := conf.Client(ctx, token)

	ytService, err := youtube.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return YouTube{}, err
	}

	return YouTube{
		service: ytService,
	}, nil

}

func (y YouTube) UpdateDescription(videoId string, desc string) error {

	ytService := y.service

	vListCall := ytService.Videos.List([]string{"snippet"})
	vListCall = vListCall.Id(videoId)
	res, err := vListCall.Do()
	if err != nil {
		return err
	}

	if len(res.Items) != 1 {
		return fmt.Errorf("wrong number of videos returned: %d", len(res.Items))
	}

	ytVideo := res.Items[0]
	ytVideo.Snippet.Description = desc

	vUpdateCall := ytService.Videos.Update([]string{"snippet"}, ytVideo)
	_, err = vUpdateCall.Do()

	return err
}
