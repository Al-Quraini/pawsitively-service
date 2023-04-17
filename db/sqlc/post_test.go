package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/alquraini/pawsitively/db/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateRandomPost(t *testing.T) Post {
	user := CreateRandomUser(t)
	image := CreateRandomImage(t)
	arg := CreatePostParams{
		Title:   sql.NullString{String: util.RandomName(), Valid: true},
		Body:    sql.NullString{String: util.RandomString(50), Valid: true},
		UserID:  user.ID,
		ImageID: uuid.NullUUID{UUID: image.ID, Valid: true},
		Status:  sql.NullString{String: "draft", Valid: true},
	}
	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.Body, post.Body)
	require.Equal(t, arg.UserID, post.UserID)
	require.Equal(t, arg.ImageID, post.ImageID)
	require.Equal(t, arg.Status, post.Status)

	require.NotZero(t, post.ID)
	require.NotZero(t, post.CreatedAt)

	return post
}

func TestCreatePost(t *testing.T) {
	CreateRandomPost(t)
}

func TestGetPost(t *testing.T) {
	post1 := CreateRandomPost(t)

	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.Title, post2.Title)
	require.Equal(t, post1.Body, post2.Body)
	require.Equal(t, post1.UserID, post2.UserID)
	require.Equal(t, post1.ImageID, post2.ImageID)
	require.Equal(t, post1.Status, post2.Status)
	require.Equal(t, post1.LikesCount, post2.LikesCount)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
}

func TestUpdatePost(t *testing.T) {
	image := CreateRandomImage(t)
	post1 := CreateRandomPost(t)

	arg := UpdatePostParams{
		ID:      post1.ID,
		Title:   sql.NullString{String: "New title", Valid: true},
		Body:    sql.NullString{String: "New body", Valid: true},
		ImageID: uuid.NullUUID{UUID: image.ID, Valid: true},
		Status:  sql.NullString{String: "new-status", Valid: true},
	}
	post2, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, arg.Title, post2.Title)
	require.Equal(t, arg.Body, post2.Body)
	require.Equal(t, post1.UserID, post2.UserID)
	require.Equal(t, arg.ImageID, post2.ImageID)
	require.Equal(t, arg.Status, post2.Status)
	require.Equal(t, post1.CreatedAt, post2.CreatedAt)
}

// func TestLikesCountOnLike(t *testing.T) {
// 	post := CreateRandomPet(t)
// 	user1 := createRandomPost
// }

func TestDeletePost(t *testing.T) {
	post := CreateRandomPost(t)

	err := testQueries.DeletePost(context.Background(), post.ID)
	require.NoError(t, err)

	post2, err := testQueries.GetPost(context.Background(), post.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, post2)
}
