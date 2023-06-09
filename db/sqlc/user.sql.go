// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    email, hashed_password
) VALUES (
    $1, $2
) RETURNING id, email, full_name, hashed_password, city, state, country, image_url, created_at, updated_at
`

type CreateUserParams struct {
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FullName,
		&i.HashedPassword,
		&i.City,
		&i.State,
		&i.Country,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, email, full_name, hashed_password, city, state, country, image_url, created_at, updated_at FROM users
WHERE 
email = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FullName,
		&i.HashedPassword,
		&i.City,
		&i.State,
		&i.Country,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, full_name, hashed_password, city, state, country, image_url, created_at, updated_at FROM users
WHERE 
id = $1
LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FullName,
		&i.HashedPassword,
		&i.City,
		&i.State,
		&i.Country,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET full_name = $2,
    image_url = $3,
    city = $4,
    state = $5,
    country = $6,
    updated_at = now()
WHERE id = $1
RETURNING id, email, full_name, hashed_password, city, state, country, image_url, created_at, updated_at
`

type UpdateUserParams struct {
	ID       int64          `json:"id"`
	FullName sql.NullString `json:"full_name"`
	ImageUrl sql.NullString `json:"image_url"`
	City     sql.NullString `json:"city"`
	State    sql.NullString `json:"state"`
	Country  sql.NullString `json:"country"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.FullName,
		arg.ImageUrl,
		arg.City,
		arg.State,
		arg.Country,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FullName,
		&i.HashedPassword,
		&i.City,
		&i.State,
		&i.Country,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
