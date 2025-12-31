-- name: InsertInscription :one
INSERT INTO inscriptions (
  competition,
  participant
) VALUES ( $1, $2 )
RETURNING *;

-- name: GetActiveCompetitionInscriptions :many
SELECT * FROM inscriptions
WHERE competition = $1 AND active = True
ORDER BY points;

-- name: GetUserInscriptions :many
SELECT * FROM inscriptions
WHERE participant = $1;
