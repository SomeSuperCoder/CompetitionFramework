-- name: InsertTaskOrder :one
INSERT INTO task_orders (
  competition,
  task
) VALUES ( $1, $2 )
RETURNING *;

-- name: GetTasksForCompetition :many
SELECT tasks.* FROM task_orders
JOIN tasks ON task_orders.task = tasks.id
WHERE competition = $1
ORDER BY task_orders.created_at DESC;

-- name: DeleteTaskOrder :one
DELETE FROM task_orders
WHERE task = $1 AND competition = $2
RETURNING *;
