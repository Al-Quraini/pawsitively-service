-- name: CreateLike :one
INSERT INTO likes (
    liked_post_id, user_id
) VALUES (
    $1, $2
) RETURNING *;


-- name: GetLikeFromPostForUser :one
SELECT * FROM likes
WHERE liked_post_id = $1
  AND user_id = $2
LIMIT 1;

-- name: GetLikesFromPost :many
SELECT * FROM likes
WHERE liked_post_id = $1;

-- name: DeleteLike :exec
DELETE FROM likes
WHERE liked_post_id = $1 AND user_id = $2;
