-- name: InsertTask :one
INSERT INTO tasks (
  name,
  details,
  points,
  duration
) VALUES ( $1, $2, $3, $4 )
RETURNING *;

-- name: FindAllTasks :many
SELECT * FROM tasks ORDER BY created_at DESC;

-- name: UpdateTask :one
UPDATE tasks
SET
    name = COALESCE(sqlc.narg('name'), name),
    details = COALESCE(sqlc.narg('details'), details),
    points = COALESCE(sqlc.narg('points'), points)
WHERE id = $1
RETURNING *;

-- name: DeleteTask :one
DELETE FROM tasks
WHERE id = $1
RETURNING *;
