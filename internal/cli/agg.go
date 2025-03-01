package cli

import (
	"context"
	"fmt"

	"github.com/cygran/gator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}
