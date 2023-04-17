package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/alquraini/pawsitively/db/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Email:          util.RandEmail(),
		HashedPassword: "1234522",
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	fmt.Print(user)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreatUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	image := CreateRandomImage(t)
	arg := UpdateUserParams{
		ID:        user1.ID,
		FirstName: sql.NullString{String: util.RandomName(), Valid: true},
		LastName:  sql.NullString{String: util.RandomName(), Valid: true},
		Email:     util.RandEmail(),
		ImageID:   uuid.NullUUID{UUID: image.ID, Valid: true},
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.FirstName, user2.FirstName)
	require.Equal(t, arg.LastName, user2.LastName)
	require.Equal(t, arg.Email, user2.Email)
	require.NotEmpty(t, user2.UpdatedAt)
}

func TestDeleteUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
