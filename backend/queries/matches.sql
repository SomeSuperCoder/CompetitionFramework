-- name: InsertMatch :one
INSERT INTO matches (
  competition,
  user1,
  user2,
  prev -- TODO: replace prev with next
) VALUES ( $1, $2, $3, $4 )
RETURNING *;

-- name: FindAllMatches :many
SELECT * FROM matches ORDER BY created_at DESC;
