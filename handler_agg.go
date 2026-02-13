package main

import (
	"context"
	"fmt"
	"time"
	"log"
	"database/sql"
	"github.com/google/uuid"
	"github.com/gcheong/gator/internal/database"
	"github.com/bjaus/gokit/dberr"
)

func handlerAgg(s *state, cmd command, user database.User) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <agg_command> time_between_requests()", cmd.Name)
	}

	timeBetweenRequests := cmd.Args[0]

	duration, err := time.ParseDuration(timeBetweenRequests)
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}

	fmt.Printf("Starting feed aggregation with a time between requests of %s...\n", duration)

	ticker := time.NewTicker(duration)

	for ;; <-ticker.C {	
		err := scrapeFeeds(context.Background(), s, user)
		if err != nil {
			return err
		}
	}

	return nil
}

func scrapeFeeds(ctx context.Context, s *state, user database.User) error {
	next_feed, err := s.db.GetNextFeedToFetch(
	ctx,
	user.ID,
	)

	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %w", err)
	}

	feed, err := fetchFeed(ctx, next_feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}	

	err = s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		UpdatedAt: time.Now().UTC(),
		ID:        next_feed.ID,
	})

	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	fmt.Printf("Fetched feed: %s\n", feed.Channel.Title)

	for _, item := range feed.Channel.Item {
		pubslishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return fmt.Errorf("error parsing publication date: %w", err)
		}
		_, err = s.db.CreatePost(ctx, database.CreatePostParams{	
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: pubslishedAt,
			FeedID:      next_feed.ID,
		})	
		if err != nil {
			if dberr.IsUniqueViolation(err) {
				continue
			}
			
			log.Printf("Error creating post for item %s: %v", item.Title, err)
		}
	}
	return nil
}
	
