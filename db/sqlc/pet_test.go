package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/alquraini/pawsitively/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomPet(t *testing.T) Pet {
	user := CreateRandomUser(t)
	arg := CreatePetParams{
		Name:             util.RandomName(),
		About:            sql.NullString{String: util.RandomString(50), Valid: true},
		UserID:           user.ID,
		Age:              int32(util.RandomAge()),
		Gender:           util.RandomGender(),
		PetType:          util.RandomAnimal(),
		Breed:            sql.NullString{String: util.RandomName(), Valid: true},
		ImageUrl:         sql.NullString{String: util.RandomImageUrl(), Valid: true},
		MedicalCondition: sql.NullString{String: util.RandomString(12), Valid: true},
	}

	pet, err := testQueries.CreatePet(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pet)

	require.Equal(t, arg.Name, pet.Name)
	require.Equal(t, arg.About, pet.About)
	require.Equal(t, arg.UserID, pet.UserID)
	require.Equal(t, arg.Age, pet.Age)
	require.Equal(t, arg.Gender, pet.Gender)
	require.Equal(t, arg.PetType, pet.PetType)
	require.Equal(t, arg.Breed, pet.Breed)
	require.Equal(t, arg.ImageUrl, pet.ImageUrl)
	require.Equal(t, arg.MedicalCondition, pet.MedicalCondition)

	require.NotZero(t, pet.ID)
	require.NotZero(t, pet.CreatedAt)

	return pet
}

func TestCreatePet(t *testing.T) {
	CreateRandomPet(t)
}

func TestGetPet(t *testing.T) {
	var lastPet Pet
	for i := 0; i < 5; i++ {
		lastPet = CreateRandomPet(t)
	}

	pets, err := testQueries.GetPets(context.Background(), lastPet.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, pets)

	for _, pet := range pets {
		require.Equal(t, pet.ID, lastPet.ID)
		require.Equal(t, pet.Name, lastPet.Name)
		require.Equal(t, pet.About, lastPet.About)
		require.Equal(t, pet.UserID, lastPet.UserID)
		require.Equal(t, pet.Age, lastPet.Age)
		require.Equal(t, pet.Gender, lastPet.Gender)
		require.Equal(t, pet.PetType, lastPet.PetType)
		require.Equal(t, pet.Breed, lastPet.Breed)
		require.Equal(t, pet.ImageUrl, lastPet.ImageUrl)
		require.Equal(t, pet.MedicalCondition, lastPet.MedicalCondition)
		require.WithinDuration(t, pet.CreatedAt, lastPet.CreatedAt, time.Second)

		require.Equal(t, pet.CreatedAt, lastPet.CreatedAt)
	}
}

func TestUpdatePet(t *testing.T) {
	pet1 := CreateRandomPet(t)

	arg := UpdatePetParams{
		Name:             util.RandomName(),
		About:            sql.NullString{String: util.RandomString(50), Valid: true},
		Age:              int32(util.RandomInt(1, 20)),
		Gender:           util.RandomGender(),
		PetType:          util.RandomAnimal(),
		Breed:            sql.NullString{String: util.RandomName(), Valid: true},
		ImageUrl:         sql.NullString{String: util.RandomImageUrl(), Valid: true},
		MedicalCondition: sql.NullString{String: util.RandomString(12), Valid: true},
		ID:               pet1.ID,
	}
	pet2, err := testQueries.UpdatePet(context.Background(), arg)
	require.NoError(t, err)
	require.NoError(t, err)
	require.NotEmpty(t, pet2)

	// ensure the updated pet has the correct information
	require.Equal(t, pet1.ID, pet2.ID)
	require.Equal(t, arg.Name, pet2.Name)
	require.Equal(t, arg.About, pet2.About)
	require.Equal(t, arg.Age, pet2.Age)
	require.Equal(t, arg.Gender, pet2.Gender)
	require.Equal(t, arg.PetType, pet2.PetType)
	require.Equal(t, arg.Breed, pet2.Breed)
	require.Equal(t, arg.ImageUrl, pet2.ImageUrl)
	require.Equal(t, pet1.CreatedAt, pet2.CreatedAt)
	require.Equal(t, pet1.CreatedAt, pet2.CreatedAt)
}

func TestDeletePet(t *testing.T) {
	pet1 := CreateRandomPet(t)
	err := testQueries.DeletePet(context.Background(), pet1.ID)
	require.NoError(t, err)

	pet2, err := testQueries.GetPetById(context.Background(), pet1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, pet2)
}
