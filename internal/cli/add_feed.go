package cli

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cygran/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	userName := s.Config.CurrentUserName
	user, err := s.Db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}
	feed, err := s.Db.AddFeed(context.Background(), database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      sql.NullString{String: name, Valid: true},
		Url:       sql.NullString{String: url, Valid: true},
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error adding feed: %w", err)
	}
	fmt.Printf("Feed added successfully: %+v\n", feed)
	return nil
}
