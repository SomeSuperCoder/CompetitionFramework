-- name: InsertCompetition :one
INSERT INTO competitions (
  name
) VALUES ( sqlc.narg('name') )
RETURNING *;

-- name: FindAllCompetitions :many
SELECT * FROM competitions ORDER BY created_at DESC;
