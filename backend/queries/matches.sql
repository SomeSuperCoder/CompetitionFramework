-- name: InsertMatch :one
INSERT INTO matches (
  name,
  competition,
  start_time,
  end_time,
  user1,
  user2,
  prev
) VALUES ( $1, $2, $3, $4, $5, $6, $7 )
RETURNING *;

-- name: FindAllMatches :many
SELECT * FROM matches ORDER BY created_at DESC;
