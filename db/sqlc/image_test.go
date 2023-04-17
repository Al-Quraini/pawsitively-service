package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/alquraini/pawsitively/db/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomImage(t *testing.T) Image {
	url := util.RandomImageUrl()
	image, err := testQueries.CreateImage(context.Background(), url)
	fmt.Print(image)
	require.NoError(t, err)
	require.NotEmpty(t, image)

	require.Equal(t, url, image.Url)

	require.NotZero(t, image.ID)

	return image
}

func TestCreateImage(t *testing.T) {
	CreateRandomImage(t)
}

func TestGetImage(t *testing.T) {
	image1 := CreateRandomImage(t)
	image2, err := testQueries.GetImage(context.Background(), image1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, image2)

	require.Equal(t, image1.ID, image2.ID)
	require.Equal(t, image1.Url, image2.Url)
}

func TestUpdateImage(t *testing.T) {
	image1 := CreateRandomImage(t)

	url := util.RandomImageUrl()
	arg := UpdateImageParams{
		ID:  image1.ID,
		Url: url,
	}
	image2, err := testQueries.UpdateImage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, image2)

	require.Equal(t, image1.ID, image2.ID)
	require.Equal(t, url, image2.Url)
}

func TestDeleteImage(t *testing.T) {
	image1 := CreateRandomImage(t)
	err := testQueries.DeleteImage(context.Background(), image1.ID)
	require.NoError(t, err)

	image2, err := testQueries.GetImage(context.Background(), image1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, image2)
}
