package aggregator

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cygran/gator/internal/database"
	"github.com/cygran/gator/internal/rss"
	"github.com/google/uuid"
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
		postID := uuid.New()
		now := time.Now().UTC()
		var publishedAt sql.NullTime
		if item.PubDate != "" {
			parsedTime, err := time.Parse(time.RFC1123Z, item.PubDate)
			if err != nil {
				parsedTime, err = time.Parse(time.RFC822, item.PubDate)
				if err != nil {
					// Continue with more formats or log the error
					log.Printf("Failed to parse date '%s': %v", item.PubDate, err)
				}
			}
			if err == nil {
				publishedAt.Time = parsedTime
				publishedAt.Valid = true
			}
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          postID,
			CreatedAt:   now,
			UpdatedAt:   now,
			Title:       sql.NullString{String: item.Title, Valid: item.Title != ""},
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "unique constraint") ||
				strings.Contains(err.Error(), "UNIQUE constraint failed") {
				continue
			}
			log.Printf("Error creating post: %v", err)
		}
	}
	return nil
}
