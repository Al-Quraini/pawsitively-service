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

func TestGetLikeFromPostForUser(t *testing.T) {
	like1 := CreateRandomLike(t)

	arg := GetLikeFromPostForUserParams{
		LikedPostID: like1.LikedPostID,
		UserID:      like1.UserID,
	}
	like2, err := testQueries.GetLikeFromPostForUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, like2)

	require.Equal(t, like1.ID, like2.ID)
	require.Equal(t, like1.LikedPostID, like2.LikedPostID)
	require.Equal(t, like1.UserID, like2.UserID)
	require.WithinDuration(t, like1.CreatedAt, like2.CreatedAt, time.Second)
}

func TestGetLikesFromPost(t *testing.T) {
	var lastLike Like
	for i := 0; i < 5; i++ {
		lastLike = CreateRandomLike(t)
	}

	likes, err := testQueries.GetLikesFromPost(context.Background(), lastLike.LikedPostID)
	require.NoError(t, err)
	require.NotEmpty(t, likes)

	for _, like := range likes {
		require.NotEmpty(t, like)
		require.Equal(t, lastLike.LikedPostID, like.LikedPostID)
	}
}

func TestDeleteLike(t *testing.T) {
	like1 := CreateRandomLike(t)

	arg1 := DeleteLikeParams{
		LikedPostID: like1.LikedPostID,
		UserID:      like1.UserID,
	}

	err := testQueries.DeleteLike(context.Background(), arg1)
	require.NoError(t, err)

	arg2 := GetLikeFromPostForUserParams{
		LikedPostID: like1.LikedPostID,
		UserID:      like1.UserID,
	}

	like2, err := testQueries.GetLikeFromPostForUser(context.Background(), arg2)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, like2)
}
