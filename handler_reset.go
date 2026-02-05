package main

// use the deletallusers query to reset the users table
import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset users table: %w", err)
	}
	fmt.Println("Users table reset successfully!")
	return nil
}
