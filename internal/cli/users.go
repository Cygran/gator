package cli

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {
	userList, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting user list: %v", err)
	}
	if len(userList) == 0 {
		return fmt.Errorf("no registered users, try the register command")
	}
	for _, user := range userList {
		if user.Name == s.Config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
