package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/alquraini/pawsitively/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomPost(t *testing.T) Post {
	user := CreateRandomUser(t)
	arg := CreatePostParams{
		Title:  sql.NullString{String: util.RandomName(), Valid: true},
		Body:   sql.NullString{String: util.RandomString(50), Valid: true},
		UserID: user.ID,
		Status: sql.NullString{String: "draft", Valid: true},
	}
	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.Body, post.Body)
	require.Equal(t, arg.UserID, post.UserID)
	require.Equal(t, arg.ImageUrl, post.ImageUrl)
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
	require.Equal(t, post1.ImageUrl, post2.ImageUrl)
	require.Equal(t, post1.Status, post2.Status)
	require.Equal(t, post1.LikesCount, post2.LikesCount)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
}

func TestListPosts(t *testing.T) {
	// var lastPost Post
	for i := 0; i < 5; i++ {
		CreateRandomPost(t)
	}

	arg := ListPostsParams{
		Limit:  5,
		Offset: 0,
	}
	posts, err := testQueries.ListPosts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, posts)

	for _, post := range posts {
		require.NotEmpty(t, post)
	}
}

func TestListPostsByUserID(t *testing.T) {
	var lastPost Post
	for i := 0; i < 5; i++ {
		lastPost = CreateRandomPost(t)
	}

	arg := ListPostsByUserIDParams{
		UserID: lastPost.UserID,
		Limit:  5,
		Offset: 0,
	}
	posts, err := testQueries.ListPostsByUserID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, posts)

	for _, post := range posts {
		require.NotEmpty(t, post)
		require.Equal(t, lastPost.UserID, post.UserID)
	}
}

func TestUpdatePost(t *testing.T) {
	post1 := CreateRandomPost(t)

	arg := UpdatePostParams{
		ID:     post1.ID,
		Title:  sql.NullString{String: "New title", Valid: true},
		Body:   sql.NullString{String: "New body", Valid: true},
		Status: sql.NullString{String: "new-status", Valid: true},
	}
	post2, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, arg.Title, post2.Title)
	require.Equal(t, arg.Body, post2.Body)
	require.Equal(t, post1.UserID, post2.UserID)
	require.Equal(t, arg.ImageUrl, post2.ImageUrl)
	require.Equal(t, arg.Status, post2.Status)
	require.Equal(t, post1.CreatedAt, post2.CreatedAt)
}

func TestDeletePost(t *testing.T) {
	post := CreateRandomPost(t)

	err := testQueries.DeletePost(context.Background(), post.ID)
	require.NoError(t, err)

	post2, err := testQueries.GetPost(context.Background(), post.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, post2)
}
