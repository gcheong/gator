package main

import (
	"context"
	"fmt"

	"github.com/gcheong/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("Usage: unfollow <feed_url>")

	}
	feedURL := cmd.Args[0]

	err := s.db.DeleteFeedByUrl(context.Background(), feedURL)
	if err != nil {
		return err
	}
	return nil
}
