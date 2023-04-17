-- name: CreatePost :one
INSERT INTO posts (
title, body, user_id, image_id, status
) VALUES (
$1, $2, $3, $4, $5
) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: UpdatePost :one
UPDATE posts
SET
title = $1,
body = $2,
image_id = $3,
status = $4,
updated_at = now()

WHERE id = $5
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;