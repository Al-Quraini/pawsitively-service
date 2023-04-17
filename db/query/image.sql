-- name: CreateImage :one
INSERT INTO images (url)
VALUES ($1)
RETURNING *;

-- name: GetImage :one
SELECT * FROM images
WHERE id = $1 LIMIT 1;

-- name: UpdateImage :one
UPDATE images
SET url = $1
WHERE id = $2
RETURNING *;

-- name: DeleteImage :exec
DELETE FROM images
WHERE id = $1;