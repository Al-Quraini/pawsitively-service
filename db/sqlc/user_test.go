package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/alquraini/pawsitively/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	password := util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	arg := CreateUserParams{
		Email:          util.RandEmail(),
		HashedPassword: hashedPassword,
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

func TestGetUserByID(t *testing.T) {
	// Create a random user
	user := CreateRandomUser(t)

	// Retrieve the user by ID
	userByID, err := testQueries.GetUserByID(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, userByID)

	// Ensure that the retrieved user is the same as the original user
	require.Equal(t, user.ID, userByID.ID)
	require.Equal(t, user.Email, userByID.Email)
	require.Equal(t, user.FullName, userByID.FullName)
	require.Equal(t, user.City, userByID.City)
	require.Equal(t, user.State, userByID.State)
	require.Equal(t, user.Country, userByID.Country)
	require.Equal(t, user.HashedPassword, userByID.HashedPassword)
	require.Equal(t, user.ImageUrl, userByID.ImageUrl)
	require.WithinDuration(t, user.CreatedAt, userByID.CreatedAt, time.Second)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	arg := UpdateUserParams{
		ID:       user1.ID,
		FullName: sql.NullString{String: util.RandomName(), Valid: true},
		City:     sql.NullString{String: util.RandomString(10), Valid: true},
		State:    sql.NullString{String: util.RandomString(10), Valid: true},
		Country:  sql.NullString{String: util.RandomString(10), Valid: true},
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.FullName, user2.FullName)
	require.Equal(t, arg.City, user2.City)
	require.Equal(t, arg.State, user2.State)
	require.Equal(t, arg.Country, user2.Country)
	require.NotEmpty(t, user2.UpdatedAt)
}

func TestDeleteUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.Email)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
