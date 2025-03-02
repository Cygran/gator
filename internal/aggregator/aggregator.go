package aggregator

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cygran/gator/internal/database"
	"github.com/cygran/gator/internal/rss"
)

func ScrapeFeeds(db *database.Queries) error {
	feed, err := db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed: %w", err)
	}
	err = db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		ID:            feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}
	rssFeed, err := rss.FetchFeed(context.Background(), feed.Url.String)
	if err != nil {
		return fmt.Errorf("error fetching feed %s: %w", feed.Name.String, err)
	}
	fmt.Printf("Feed: %s\n", feed.Name.String)
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("- %s\n", item.Title)
	}

	return nil
}
