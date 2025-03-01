package cli

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cygran/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username required")
	}
	userName := cmd.Args[0]
	_, err := s.Db.GetUser(context.Background(), userName)
	if err != nil {
		if err == sql.ErrNoRows {
			user, err := s.Db.CreateUser(
				context.Background(),
				database.CreateUserParams{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					Name:      userName,
				},
			)
			if err != nil {
				return err
			}
			s.Config.SetUser(user.Name)
			fmt.Printf("User %s created successfully!\nUser details: %+v\n", userName, user)
			return nil
		} else {
			// Something went wrong (e.g., database issue).
			return fmt.Errorf("error fetching user: %w", err)
		}
	} else {
		// The user already exists, so exit with an error.
		return fmt.Errorf("user %s already exists", userName)
	}
}
