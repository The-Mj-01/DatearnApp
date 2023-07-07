package image

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestImageService_GetAllImage(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createImageService(db)

	_, err = sv.GetAllImage(nil, nil, nil, nil, nil, 0)
	assert.Error(t, err, "Expected interest not found error")
	assert.ErrorIs(t, err, ImageNotFound, "Expected interest not found error")
}

func TestImageService_CreateImage(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createImageService(db)

	width := 200
	height := 200
	randImageableId := uint(rand.Int())
	randImageableType := "Bio"
	randImageName := "img1.jpg"
	img := mockImage(width, height, randImageableId, randImageableType, randImageName)

	createdImage, err := sv.CreateImage(img.ImageableId, img.Name, img.Path, img.ImageableType)
	defer destructCreatedObjects(db, []Image{*createdImage})

	assert.NoError(t, err, "image service bio creation failed")
	assert.Equal(t, img.Name, createdImage.Name)
}

func createImageService(db *gorm.DB) ImageServiceInterface {
	return NewImageService(NewImageRepository(db))
}
