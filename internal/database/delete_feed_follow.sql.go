// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: delete_feed_follow.sql

package database

import (
	"context"
	"database/sql"
)

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
USING feeds, users
WHERE feed_follows.feed_id = feeds.id
AND feed_follows.user_id = users.id
AND feeds.url = $1
AND users.name = $2
`

type DeleteFeedFollowParams struct {
	Url  sql.NullString
	Name string
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, arg.Url, arg.Name)
	return err
}
