// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
	"time"
)

type Like struct {
	ID          int64     `json:"id"`
	LikedPostID int64     `json:"liked_post_id"`
	UserID      int64     `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type Pet struct {
	ID               int64          `json:"id"`
	Name             string         `json:"name"`
	About            sql.NullString `json:"about"`
	UserID           int64          `json:"user_id"`
	Age              int32          `json:"age"`
	Gender           string         `json:"gender"`
	PetType          string         `json:"pet_type"`
	Breed            sql.NullString `json:"breed"`
	ImageUrl         sql.NullString `json:"image_url"`
	MedicalCondition sql.NullString `json:"medical_condition"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        sql.NullTime   `json:"updated_at"`
}

type Post struct {
	ID         int64          `json:"id"`
	Title      sql.NullString `json:"title"`
	Body       sql.NullString `json:"body"`
	UserID     int64          `json:"user_id"`
	ImageUrl   sql.NullString `json:"image_url"`
	Status     sql.NullString `json:"status"`
	LikesCount int32          `json:"likes_count"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  sql.NullTime   `json:"updated_at"`
}

type User struct {
	ID             int64          `json:"id"`
	Email          string         `json:"email"`
	FullName       sql.NullString `json:"full_name"`
	HashedPassword string         `json:"hashed_password"`
	City           sql.NullString `json:"city"`
	State          sql.NullString `json:"state"`
	Country        sql.NullString `json:"country"`
	ImageUrl       sql.NullString `json:"image_url"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      sql.NullTime   `json:"updated_at"`
}
