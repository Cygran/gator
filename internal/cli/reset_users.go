package cli

import (
	"context"
	"fmt"
)

func HandlerResetUsers(s *State, cmd Command) error {
	fmt.Println("Resetting users database...")
	err := s.Db.Reset(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Database reset successfully")
	return nil
}
