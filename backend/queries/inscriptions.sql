-- name: InsertInscription :one
INSERT INTO inscriptions (
  competition,
  participant
) VALUES ( $1, $2 )
RETURNING *;

-- name: GetCompetitionInscriptions :many
SELECT i.*, u.id AS user_id
FROM inscriptions i
JOIN users u ON i.participant = u.id
WHERE i.competition = $1
ORDER BY i.created_at ASC;

-- name: GetUserInscriptions :many
SELECT * FROM inscriptions
WHERE participant = $1;
