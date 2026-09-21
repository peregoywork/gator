package main

import (
	"fmt"
	"time"
	"context"
	"database/sql"

	"github.com/google/uuid"

	"gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("agg function must be given some duration (eg: 1s, 1m, 1h)")
	}

	dur, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("could not parse given arg as time.Duration: %s", cmd.Args[0])
	}
	
	ticker := time.NewTicker(dur)
	fmt.Println("Collecting feeds every %s", dur)

	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("error scraping feeds: %s", err)
		}
	}

	return nil
}


func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("could not get next feeds to fetch: %s", err) 
	}

	fmt.Printf("scraping next feed: %s", feed.Name)

	scrapeFeed(s.db, feed)

	return nil
}


func scrapeFeed(db *database.Queries, feed database.Feed) error {
	err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		saveRssItemAsPost(db, item, feed.ID)
	}

	return nil
}


func saveRssItemAsPost(db *database.Queries, item RSSItem, feed_id uuid.UUID) {
	publishedAt := sql.NullTime{}
	if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
		publishedAt = sql.NullTime{
			Time: t,
			Valid: true,
		}
	} 

	post, err := db.CreatePost(context.Background(), database.CreatePostParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title: item.Title,
		Url: item.Link,
		Description: sql.NullString{ String: item.Description, Valid: true },
		PublishedAt: publishedAt,
		FeedID: feed_id,
	})

	if err != nil {
		// errors: ignore if url already exists, log others
		fmt.Printf("error while creating post: %v\n", err)
		// unique constraint "posts_url_key"
		return
	}

	fmt.Printf("Post Saved: %v\n", post)
}

