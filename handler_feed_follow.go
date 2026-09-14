package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"

	"gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("follow command requires one args: url")
	}

	url := cmd.Args[0]
	ctx := context.Background()

	user, err := s.db.GetOneUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error fetching user: %w", err)
	}

	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("could not find feed for given url: %w", err)
	}

	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not follow feed: %w", err)
	}

	fmt.Printf("%s is now following feed %s", user.Name, feed.Name)

	return nil
}

func handlerListFollowing(s *state, cmd command) error {
	ctx := context.Background()
	follows, err := s.db.GetFeedFollowsForUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could not list all feeds: %w", err)
	}

	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}

