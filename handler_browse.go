package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gcheong/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	
	var post_limit int = 2
	
	if len(cmd.Args) != 0 {
		limit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit argument: %w", err)
		}
		post_limit = limit
	}

	posts, err := s.db.GetPostsForUser(context.Background(), user.ID)

	if err != nil {
		return err
	}

	post_count := 0

	for _, post := range posts {
		fmt.Printf("Title: %s\nLink: %s\nPublished: %s\n\n", post.Title, post.Url, post.PublishedAt)
		post_count++
		if post_count >= post_limit {
			break
		}
	}

	return nil
}

