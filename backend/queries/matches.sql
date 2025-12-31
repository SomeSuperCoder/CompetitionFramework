-- name: InsertMatch :one
INSERT INTO matches (
  name,
  competition,
  start_time,
  end_time
) VALUES ( $1, $2, $3, $4 )
RETURNING *;

-- name: FindAllMatches :many
SELECT * FROM matches ORDER BY created_at DESC;
