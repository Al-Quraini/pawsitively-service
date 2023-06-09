// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: post.sql

package db

import (
	"context"
	"database/sql"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
title, body, user_id, image_url, status
) VALUES (
$1, $2, $3, $4, $5
) RETURNING id, title, body, user_id, image_url, status, likes_count, created_at, updated_at
`

type CreatePostParams struct {
	Title    sql.NullString `json:"title"`
	Body     sql.NullString `json:"body"`
	UserID   int64          `json:"user_id"`
	ImageUrl sql.NullString `json:"image_url"`
	Status   sql.NullString `json:"status"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.Title,
		arg.Body,
		arg.UserID,
		arg.ImageUrl,
		arg.Status,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.ImageUrl,
		&i.Status,
		&i.LikesCount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getPost = `-- name: GetPost :one
SELECT id, title, body, user_id, image_url, status, likes_count, created_at, updated_at FROM posts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.ImageUrl,
		&i.Status,
		&i.LikesCount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listPosts = `-- name: ListPosts :many
SELECT id, title, body, user_id, image_url, status, likes_count, created_at, updated_at FROM posts
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPostsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPosts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.UserID,
			&i.ImageUrl,
			&i.Status,
			&i.LikesCount,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPostsByUserID = `-- name: ListPostsByUserID :many
SELECT id, title, body, user_id, image_url, status, likes_count, created_at, updated_at FROM posts
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListPostsByUserIDParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPostsByUserID(ctx context.Context, arg ListPostsByUserIDParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPostsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.UserID,
			&i.ImageUrl,
			&i.Status,
			&i.LikesCount,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET
title = $1,
body = $2,
image_url = $3,
status = $4,
updated_at = now()

WHERE id = $5
RETURNING id, title, body, user_id, image_url, status, likes_count, created_at, updated_at
`

type UpdatePostParams struct {
	Title    sql.NullString `json:"title"`
	Body     sql.NullString `json:"body"`
	ImageUrl sql.NullString `json:"image_url"`
	Status   sql.NullString `json:"status"`
	ID       int64          `json:"id"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost,
		arg.Title,
		arg.Body,
		arg.ImageUrl,
		arg.Status,
		arg.ID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.ImageUrl,
		&i.Status,
		&i.LikesCount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
