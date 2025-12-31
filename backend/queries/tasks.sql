-- name: InsertTask :one
INSERT INTO tasks (
  name,
  details
) VALUES ( $1, $2 )
RETURNING *;

-- name: FindAllTasks :many
SELECT * FROM tasks ORDER BY created_at DESC;
