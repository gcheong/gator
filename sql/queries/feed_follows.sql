-- name: CreateFeedFollow :one
with inserted_feed_follow as (
    insert into feed_follows (id, created_at, updated_at, user_id, feed_id) 
    values ($1, $2, $3, $4, $5)
    returning *
)
SELECT inserted_feed_follow.*,
feeds.name as feed_name,
users.name as user_name
FROM inserted_feed_follow
JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
JOIN users ON inserted_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT ff.*, f.name as feed_name, u.name as user_name
FROM feed_follows ff
JOIN feeds f ON ff.feed_id = f.id
JOIN users u ON ff.user_id = u.id
WHERE ff.user_id = $1;