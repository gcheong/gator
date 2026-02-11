-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.name, f.url, u.name as created_by FROM feeds f JOIN users u ON f.user_id = u.id;

-- name: GetFeedByUrl :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeedFollowsByUserId :many
SELECT ff.*, f.name as feed_name, u.name as user_name
FROM feed_follows ff
JOIN feeds f ON ff.feed_id = f.id
JOIN users u ON ff.user_id = u.id
WHERE ff.user_id = $1;