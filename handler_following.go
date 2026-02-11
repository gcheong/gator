package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {

	curent_user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), curent_user.ID)

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("Feed Name: %s, followed by: %s\n", feed.FeedName, feed.UserName)
	}

	return nil

}
