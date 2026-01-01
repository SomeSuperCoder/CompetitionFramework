-- name: InsertMatch :one
INSERT INTO matches (
  competition,
  user1,
  user2,
  next
) VALUES ( $1, $2, $3, $4 )
RETURNING *;

-- name: FindAllMatches :many
SELECT * FROM matches ORDER BY created_at DESC;

-- name: FindAllRunningMatches :many
SELECT * FROM competitions WHERE status = 'running' ORDER BY created_at ASC;
