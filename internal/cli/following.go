package cli

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {
	user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}
	followList, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get followed feed list for %v", user.Name)
	}
	if len(followList) == 0 {
		return fmt.Errorf("user %v is not following any feeds. Try the follow or addfeed commands", user.Name)
	}
	for i, follow := range followList {
		fmt.Printf("%v. %v", i+1, follow.FeedName.String)
	}
	return nil
}
