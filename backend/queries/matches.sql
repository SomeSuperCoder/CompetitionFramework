-- name: InsertMatch :one
INSERT INTO matches (
  competition,
  user1,
  user2,
  next
) VALUES ( $1, $2, $3, $4 )
RETURNING *;

-- name: FindAllMatches :many
SELECT * FROM matches ORDER BY created_at DESC;

-- name: FindAllRunningMatchesInCompetition :many
SELECT * FROM matches WHERE status = 'running' AND competition = $1 ORDER BY created_at ASC;

-- name: FindAllCompletedMatchesInCompetitionWithNoDescendents :many
SELECT * FROM matches WHERE status = 'completed' AND competition = $1 AND next IS NULL ORDER BY created_at ASC;

-- name: FindAllRunningMatches :many
SELECT * FROM matches WHERE status = 'running' ORDER BY created_at ASC;

-- name: SetNextForMatch :one
UPDATE matches
SET next = $2
WHERE id = $1
RETURNING *;
