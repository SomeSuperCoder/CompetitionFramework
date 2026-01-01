-- name: InsertCompetition :one
INSERT INTO competitions (
  name
) VALUES ( $1 )
RETURNING *;

-- name: FindAllCompetitions :many
SELECT * FROM competitions ORDER BY created_at DESC;

-- name: FindAllRunningCompetitions :many
SELECT * FROM competitions WHERE status = 'running' ORDER BY created_at ASC;

-- name: RenameCompetition :one
UPDATE competitions
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCompetition :one
DELETE FROM competitions
WHERE id = $1
RETURNING *;
