package cli

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cygran/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *State, cmd Command) error {
	user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}
	if len(cmd.Args) == 0 {
		return fmt.Errorf("url required")
	}
	url := sql.NullString{String: cmd.Args[0], Valid: true}
	feed, err := s.Db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("feed not found: %v. Please use 'addfeed' command to add this feed first", url)
	}
	follow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to follow feed: %v", err)
	}
	fmt.Printf("Feed: %v, followed by %v\n", follow.FeedName, follow.UserName)
	return nil
}
