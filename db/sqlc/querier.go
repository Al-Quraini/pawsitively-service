// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateLike(ctx context.Context, arg CreateLikeParams) (Like, error)
	CreatePet(ctx context.Context, arg CreatePetParams) (Pet, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteLike(ctx context.Context, arg DeleteLikeParams) error
	DeletePet(ctx context.Context, id int64) error
	DeletePost(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetLikeFromPostForUser(ctx context.Context, arg GetLikeFromPostForUserParams) (Like, error)
	GetLikesFromPost(ctx context.Context, likedPostID int64) ([]Like, error)
	GetPetById(ctx context.Context, id int64) (Pet, error)
	GetPets(ctx context.Context, userID int64) ([]Pet, error)
	GetPost(ctx context.Context, id int64) (Post, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUser(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error)
	ListPostsByUserID(ctx context.Context, arg ListPostsByUserIDParams) ([]Post, error)
	UpdatePet(ctx context.Context, arg UpdatePetParams) (Pet, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
