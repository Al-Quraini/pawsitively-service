package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomLike(t *testing.T) Like {
	post := CreateRandomPost(t)
	user := CreateRandomUser(t)

	like, err := testQueries.CreateLike(context.Background(), CreateLikeParams{
		LikedPostID: post.ID,
		UserID:      user.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, like)

	require.Equal(t, post.ID, like.LikedPostID)
	require.Equal(t, user.ID, like.UserID)

	return like
}

func TestCreateLike(t *testing.T) {
	CreateRandomLike(t)
}

func TestGetLike(t *testing.T) {
	like1 := CreateRandomLike(t)
	like2, err := testQueries.GetLike(context.Background(), like1.LikedPostID)
	require.NoError(t, err)
	require.NotEmpty(t, like2)

	require.Equal(t, like1.ID, like2.ID)
	require.Equal(t, like1.LikedPostID, like2.LikedPostID)
	require.Equal(t, like1.UserID, like2.UserID)
	require.WithinDuration(t, like1.CreatedAt, like2.CreatedAt, time.Second)
}

func TestDeleteLike(t *testing.T) {
	like1 := CreateRandomLike(t)
	err := testQueries.DeleteLike(context.Background(), like1.ID)
	require.NoError(t, err)

	like2, err := testQueries.GetLike(context.Background(), like1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, like2)
}
