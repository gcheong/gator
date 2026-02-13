package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gcheong/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <feed_name> <feed_url>", cmd.Name)
	}
	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      feedName,
		Url:       feedURL,
	})

	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Printf("Feed added successfully! %v\n", feed)

	follow_cmd := command{
		Name: "follow",
		Args: []string{feedURL},
	}

	err = handlerFollow(s, follow_cmd, user)

	if err != nil {
		return fmt.Errorf("couldn't follow feed after adding: %w", err)
	}

	return nil
}
