package cli

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {
	feedList, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feed list: %v", err)
	}
	if len(feedList) == 0 {
		return fmt.Errorf("no registered feeds, try the addfeed command")
	}

	for i, feed := range feedList {
		user, err := s.Db.GetUserByID(context.Background(), feed.UserID)
		name := feed.Name.String
		url := feed.Url.String
		if err != nil {
			return fmt.Errorf("error getting username that added feed %v: %v", feed.Name, err)
		}
		fmt.Printf("Feed %v:\nName: %v,\nURL: %v,\nUser: %v\n\n", i+1, name, url, user.Name)
	}
	return nil
}
