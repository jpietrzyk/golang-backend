-- name: CreateGame :one
INSERT INTO games (
  owner_id,
  starts_at,
  ends_at
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetGame :one
SELECT * FROM games
WHERE id = $1 LIMIT 1;

-- name: ListGames :many
SELECT * FROM games
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateGame :one
UPDATE games
SET
  starts_at = $2,
  ends_at = $3,
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteGame :exec
DELETE FROM games
WHERE id = $1;
