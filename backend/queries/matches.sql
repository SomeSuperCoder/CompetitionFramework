-- name: InsertMatch :one
INSERT INTO matches (
  competition,
  user1,
  user2,
  next
) VALUES ( $1, $2, $3, $4 )
RETURNING *;

-- name: FindAllLeafMatchesOfCompetiton :many
SELECT * FROM matches
WHERE competition = $1 AND
      next IS NULL AND
      status = "completed";

-- name: FindAllMatches :many
SELECT * FROM matches ORDER BY created_at DESC;

-- name: FindAllRunningMatchesInCompetition :many
SELECT * FROM matches WHERE status = 'running' AND competition = $1 ORDER BY created_at ASC;

-- name: GetCompetitionDescendentlessMatchStats :many
SELECT
  competition,
  COUNT(CASE WHEN status = 'completed' THEN 1 END) AS completed_count,
  COUNT(*) AS total_count
FROM matches
WHERE next IS NULL
GROUP BY competition
ORDER BY created_at ASC;

-- name: GetLockedMatchRoundStats :many
SELECT
  matches.id as match,
  COUNT(CASE WHEN rounds.status = 'completed' THEN 1 END) AS completed_count,
  competitions.min_rounds,
  matches.user1_points,
  matches.user2_points
FROM matches
JOIN competitions ON competition.id = matches.competition
LEFT JOIN rounds ON rounds.match = matches.id
WHERE matches.status = 'running'
GROUP BY matches.id, competitions.min_rounds, matches.user1_points, matches.user2_points
HAVING COUNT(CASE WHEN rounds.status = 'running' THEN 1 END) = 0;

-- name: SetNextForMatch :one
UPDATE matches
SET next = $2
WHERE id = $1
RETURNING *;
