package cli

import (
	"context"
	"database/sql"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username required")
	}
	userName := cmd.Args[0]
	user, err := s.Db.GetUser(context.Background(), userName)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user does not exist try the register command")
		} else {
			return err
		}
	}
	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("Current user set to: %s\n", user.Name)
	return nil
}
