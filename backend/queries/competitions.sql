-- name: FinishCompetition :exec
WITH final_match AS (
  SELECT winner
  FROM matches
  WHERE competition = $1
  AND next IS NULL
  LIMIT 1
)
UPDATE competitions
SET
  winner = (SELECT winner FROM final_match),
  status = 'completed'
WHERE competitions.id = $1;

-- name: InsertCompetition :one
INSERT INTO competitions (
  name,
  start_time,
  min_rounds
) VALUES ( $1, $2, $3 )
RETURNING *;

-- name: FindAllCompetitions :many
SELECT * FROM competitions ORDER BY created_at DESC;

-- name: FindAllRunningCompetitions :many
SELECT * FROM competitions WHERE status = 'running' ORDER BY created_at ASC;

-- name: FindAllCompetitionsToStart :many
SELECT * FROM competitions
WHERE
  status = 'awaiting' AND
  start_time < CURRENT_TIMESTAMP 
ORDER BY start_time ASC;

-- name: SetCompetitionStatus :one
UPDATE competitions
SET
  status = $2
WHERE id = $1
RETURNING *;

-- name: RenameCompetition :one
UPDATE competitions
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCompetition :one
DELETE FROM competitions
WHERE id = $1
RETURNING *;
