package cli

import (
	"fmt"
	"time"

	"github.com/cygran/gator/internal/aggregator"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("require a time between requests")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to parse request frequency: %v", err)
	}
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		aggregator.ScrapeFeeds(s.Db)
	}
}
