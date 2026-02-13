package main

import (
	"context"
	"fmt"

	"github.com/gcheong/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("Feed Name: %s, followed by: %s\n", feed.FeedName, feed.UserName)
	}

	return nil

}
