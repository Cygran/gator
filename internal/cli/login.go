package cli

import "fmt"

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username required")
	}
	userName := cmd.Args[0]
	err := s.Config.SetUser(userName)
	if err != nil {
		return err
	}
	fmt.Printf("Current user set to: %s\n", userName)
	return nil
}
