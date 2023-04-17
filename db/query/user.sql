-- name: CreateUser :one
INSERT INTO users (
    email, hashed_password
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET first_name = $2,
    last_name = $3,
    email = $4,
    image_id = $5,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec 
DELETE FROM users
WHERE id = $1;