package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/d-darac/gator/internal/database"
)

func handler_browse(s *state, cmd command, user database.User) error {
	limit := 2

	if len(cmd.Args) > 0 {
		i, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("limit argument not numeric\nexample usage: browse 5")
		}
		limit = i
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	fmt.Printf("found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon 2 Jan"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}
