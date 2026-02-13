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

-- name: MarkFeedFetched :exec
UPDATE feeds SET updated_at = $1, last_fetched_at = $1 WHERE id = $2;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds WHERE user_id = $1 ORDER BY last_fetched_at ASC NULLS FIRST LIMIT 1;

-- name: DeleteFeedByUrl :exec
DELETE FROM feeds WHERE url = $1;