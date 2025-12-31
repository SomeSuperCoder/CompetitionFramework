-- name: InsertTaskOrder :one
INSERT INTO task_orders (
  competition,
  task
) VALUES ( $1, $2 )
RETURNING *;

-- name: GetTasksForCompetition :many
SELECT * FROM task_orders
WHERE competition = $1
ORDER BY created_at DESC;
