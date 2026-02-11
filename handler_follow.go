package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gcheong/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}
	feedURL := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), feedURL)

	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Printf("Feed: %s followed successfully! by user: %s\n", feed_follow.FeedName, feed_follow.UserName)

	return nil
}
