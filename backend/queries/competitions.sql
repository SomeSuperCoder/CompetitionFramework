-- name: InsertCompetition :one
INSERT INTO competitions (
  name
) VALUES ( $1 )
RETURNING *;

-- name: FindAllCompetitions :many
SELECT * FROM competitions ORDER BY created_at DESC;
