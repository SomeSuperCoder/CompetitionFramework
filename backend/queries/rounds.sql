-- name: InsertRound :one
INSERT INTO rounds (
  task,
  match,
  start_time,
  end_time
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: FindAllCompletedRoundsInMatch :many
SELECT * FROM rounds WHERE status = 'completed' AND match = $1 ORDER BY created_at ASC;

-- name: CreateNewRoundForMatch :one
WITH match AS (
  SELECT competition FROM matches
  WHERE id = $1
),
task_data AS (
  SELECT
    tasks.id as task_id,
    tasks.duration as task_duration
  FROM task_orders
  JOIN tasks ON task_orders.task = tasks.id
  WHERE task_orders.competition = (SELECT competition FROM match)
  ORDER BY RANDOM()
  LIMIT 1
)
INSERT INTO rounds (
  task,
  match,
  start_time,
  end_time
) VALUES (
  (SELECT task_id FROM task_data),
  $1,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP + (SELECT task_duration FROM task_data)
)
RETURNING *;
