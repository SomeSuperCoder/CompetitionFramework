-- name: InsertUser :one
INSERT INTO users (
  name,
  email,
  password
) VALUES (
  $1,
  $2,
  crypt($3, gen_salt('bf'))
) RETURNING 
  id,
  name,
  email,
  created_at,
  role
;

-- name: UpdateProfile :one
UPDATE users SET
  name = COALESCE(sqlc.narg('name'), name)
WHERE id = sqlc.arg('id')
RETURNING name, 1 as _dummy;
