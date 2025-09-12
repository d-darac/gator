package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func handlerAggregate(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't parse duration: %w", err)
	}

	log.Printf("collecting feeds every %v", timeBetweenReqs.String())

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	nextToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	log.Println("found a feed to fetch")

	err = s.db.MarkFeedFetched(context.Background(), nextToFetch.ID)
	if err != nil {
		return fmt.Errorf("couldn't update feed: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), nextToFetch.Url)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("      Found post: %s\n", item.Title)
	}
	log.Printf("feed %s collected, %v posts found", nextToFetch.Name, len(rssFeed.Channel.Item))

	return nil
}
