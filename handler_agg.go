package main

import (
	"fmt"
	"time"

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
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return log.Errorf("could not get next feeds to fetch: %s", err) 
	}

	scrapeFeed(s.db, feed)

	return nil
}


func scrapeFeed(db *database.Queries, feed database.Feed) {
	err = db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}

	fmt.Println("Feed:")
	fmt.Printf("%s - %s - %s\n", rssFeed.Channel.Title, rssFeed.Channel.Description, rssFeed.Channel.Link)
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("%+v\n", item.Title)
	}
}


