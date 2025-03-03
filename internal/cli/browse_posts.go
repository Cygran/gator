package cli

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/cygran/gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 0 {
		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}

		if parsedLimit < 1 {
			return fmt.Errorf("limit must be a positive number")
		}

		limit = parsedLimit
	}
	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error fetching posts: %w", err)
	}
	if len(posts) == 0 {
		fmt.Println("No posts found. Try following some feeds first!")
		return nil
	}
	fmt.Printf("Showing %d most recent posts:\n\n", len(posts))

	for i, post := range posts {
		fmt.Printf("--- Post %d ---\n", i+1)
		fmt.Printf("Title: %s\n", post.Title.String)
		fmt.Printf("URL: %s\n", post.Url)

		if post.PublishedAt.Valid {
			fmt.Printf("Published: %s\n", post.PublishedAt.Time.Format(time.RFC1123))
		}
		if post.Description.Valid && post.Description.String != "" {
			// Truncate description if it's too long for display
			desc := post.Description.String
			if len(desc) > 200 {
				desc = desc[:200] + "..."
			}
			fmt.Printf("Description: %s\n", desc)
		}

		fmt.Printf("\n")
	}
	return nil
}
