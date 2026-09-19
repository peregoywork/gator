-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id, last_updated)
VALUES(
    $1, 
    $2, 
    $3, 
    $4, 
    $5, 
    $6,
    $7
)
RETURNING *;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: MarkFeedFetched :exec
UPDATE feeds SET updated_at = NOW(), last_updated = NOW() WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds ORDER BY last_updated NULLS FIRST;

