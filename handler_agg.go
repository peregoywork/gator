package main

import (
	"fmt"
	"context"
)

func handlerAgg(s *state, cmd command) error {
	ctx := context.Background()
	url := "https://www.wagslane.dev/index.xml"
	rss, err := fetchFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("Could not fetch feed: %s\n", err)
	}

	fmt.Println("Feed:")
	fmt.Printf("%s - %s - %s\n", rss.Channel.Title, rss.Channel.Description, rss.Channel.Link)
	for _, item := range rss.Channel.Item {
		fmt.Printf("%+v\n", item)
	}

	return nil
}

