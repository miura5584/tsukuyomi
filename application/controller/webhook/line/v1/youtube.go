package v1

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"os"
)

func FetchLatestVideo() *youtube.SearchResult {
	apiKey := os.Getenv("YOUTUBE_API_KEY")

	youtubeService, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		// TODO:
	}

	youtubeSearchList := youtubeService.Search.List([]string{"snippet"}).MaxResults(5)
	res, err := youtubeSearchList.Do()
	if err != nil {
		// TODO:
	}

	return res.Items[0]
}