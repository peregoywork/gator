package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"

	"gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Feed requires two args: name, url")
	}

	feed_name := cmd.Args[0]
	feed_url := cmd.Args[1]
	ctx := context.Background()

	user, err := s.db.GetOneUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user not found for name %s. Cannot create feed", s.cfg.CurrentUserName)
	}

	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: feed_name,
		Url: feed_url,
		UserID: user.ID,
	})

	if err != nil {
		return fmt.Errorf("Failed to create feed: %s", err)
	}

	fmt.Printf("%+v\n", feed)
	
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("Failed to fetch list of feeds")
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(ctx, feed.UserID)
		if err != nil {
			fmt.Errorf("could not get user: %w" , err)
		} 
		fmt.Printf("%s %s %s\n", feed.Name, feed.Url, user.Name)
	}

	return nil
}
