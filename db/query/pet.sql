-- name: CreatePet :one
INSERT INTO pets (
    name, about, user_id, age, gender, pet_type, breed, image_id, medical_condition
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetPet :one
SELECT * FROM pets
WHERE id = $1 LIMIT 1;

-- name: UpdatePet :one
UPDATE pets
SET 
    name = $1,
    about = $2,
    age = $3,
    gender = $4,
    pet_type = $5,
    breed = $6,
    image_id = $7,
    medical_condition = $8,
    updated_at = now()

WHERE id = $9
RETURNING *;

-- name: DeletePet :exec 
DELETE FROM pets
WHERE id = $1;