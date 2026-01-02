-- name: InsertRound :one
INSERT INTO rounds (
  task,
  match,
  start_time,
  end_time
) VALUES (
  $1, $2, $3, $4
) RETURNING *;
