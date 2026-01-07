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
