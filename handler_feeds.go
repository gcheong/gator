package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed Name: %s, URL: %s, Created By: %s\n", feed.Name, feed.Url, feed.CreatedBy)
	}

	return nil
}
