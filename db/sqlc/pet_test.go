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

func CreateRandomPet(t *testing.T) Pet {
	user := CreateRandomUser(t)
	image := CreateRandomImage(t)
	arg := CreatePetParams{
		Name:             util.RandomName(),
		About:            sql.NullString{String: util.RandomString(50), Valid: true},
		UserID:           user.ID,
		Age:              int32(util.RandomAge()),
		Gender:           util.RandomGender(),
		PetType:          util.RandomAnimal(),
		Breed:            sql.NullString{String: util.RandomName(), Valid: true},
		ImageID:          uuid.NullUUID{UUID: image.ID, Valid: true},
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
	require.Equal(t, arg.ImageID, pet.ImageID)
	require.Equal(t, arg.MedicalCondition, pet.MedicalCondition)

	require.NotZero(t, pet.ID)
	require.NotZero(t, pet.CreatedAt)

	return pet
}

func TestCreatePet(t *testing.T) {
	CreateRandomPet(t)
}

func TestGetPet(t *testing.T) {
	pet1 := CreateRandomPet(t)

	pet2, err := testQueries.GetPet(context.Background(), pet1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pet2)

	require.Equal(t, pet1.ID, pet2.ID)
	require.Equal(t, pet1.Name, pet2.Name)
	require.Equal(t, pet1.About, pet2.About)
	require.Equal(t, pet1.UserID, pet2.UserID)
	require.Equal(t, pet1.Age, pet2.Age)
	require.Equal(t, pet1.Gender, pet2.Gender)
	require.Equal(t, pet1.PetType, pet2.PetType)
	require.Equal(t, pet1.Breed, pet2.Breed)
	require.Equal(t, pet1.ImageID, pet2.ImageID)
	require.Equal(t, pet1.MedicalCondition, pet2.MedicalCondition)
	require.WithinDuration(t, pet1.CreatedAt, pet2.CreatedAt, time.Second)

	require.Equal(t, pet1.CreatedAt, pet2.CreatedAt)
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
		ImageID:          uuid.NullUUID{UUID: pet1.ImageID.UUID, Valid: true},
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
	require.Equal(t, arg.ImageID, pet2.ImageID)
	require.Equal(t, pet1.CreatedAt, pet2.CreatedAt)
	require.Equal(t, pet1.CreatedAt, pet2.CreatedAt)
}

func TestDeletePet(t *testing.T) {
	pet1 := CreateRandomPet(t)
	err := testQueries.DeletePet(context.Background(), pet1.ID)
	require.NoError(t, err)

	pet2, err := testQueries.GetPet(context.Background(), pet1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, pet2)
}
