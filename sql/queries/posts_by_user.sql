-- name: GetPostsForUser :many
SELECT feeds.name, posts.title, posts.url, posts.description, posts.published_at FROM posts
INNER JOIN feeds
ON posts.feed_id = feeds.id
WHERE feeds.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;