package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {

	var feedURL = "https://wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Printf("%v\n", *feed)

	return nil
}
