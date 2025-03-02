package cli

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cygran/gator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("feed url required")
	}
	url := sql.NullString{String: cmd.Args[0], Valid: true}
	err := s.Db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		Url:  url,
		Name: user.Name,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %v", err)
	}
	return nil
}
