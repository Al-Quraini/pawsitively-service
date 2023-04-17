-- name: CreateLike :one
INSERT INTO likes (
    liked_post_id, user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetLike :one
SELECT * FROM likes
WHERE liked_post_id = $1 LIMIT 1;

-- name: DeleteLike :exec
DELETE FROM likes
WHERE id = $1;
