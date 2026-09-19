package main

import (
	"fmt"
	"time"
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


