-- name: CreateUser :one
INSERT INTO users (
    email, hashed_password
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE 
id = $1
LIMIT 1;

-- name: GetUser :one
SELECT * FROM users
WHERE 
email = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET full_name = $2,
    image_url = $3,
    city = $4,
    state = $5,
    country = $6,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec 
DELETE FROM users
WHERE id = $1;