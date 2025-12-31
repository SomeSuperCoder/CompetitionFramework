-- name: InsertTask :one
INSERT INTO tasks (
  name,
  details,
  competition
) VALUES ( $1, $2, $3 )
RETURNING *;

-- name: FindAllTasks :many
SELECT * FROM tasks ORDER BY created_at DESC;
