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

	width := []int{200, 200, 200, 200, 200}
	height := []int{200, 200, 200, 200, 200}
	randImageableId := []uint{uint(rand.Int()), uint(rand.Int()), uint(rand.Int()), uint(rand.Int()), uint(rand.Int())}
	randImageableType := []string{"Bio", "Bio", "Bio", "Bio", "Bio"}
	randImageName := []string{"img1.jpg", "img2.jpg", "img3.jpg", "img4.jpg", "img5.jpg"}

	mockedImage := mockAndInsertImage(db, width, height, randImageableId, randImageableType, randImageName, 5)
	defer destructCreatedObjects(db, mockedImage)
	defer removeCreatedImageFile(mockedImage)

	fetchedImage, err := sv.GetAllImage(&mockedImage[0].Id, &mockedImage[0].ImageableId, &mockedImage[0].Name, &mockedImage[0].ImageableType, nil, 0)

	assert.NoError(t, err, "Get all Images fetch query failed")

	assertImage(t, mockedImage, *fetchedImage)
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
	defer removeCreatedImageFile([]Image{*img})

	assert.NoError(t, err, "image service bio creation failed")
	assert.Equal(t, img.Name, createdImage.Name)
}

func TestImageService_UpdateImage(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createImageService(db)

	width := []int{200, 400}
	height := []int{200, 400}
	randImageableId := []uint{uint(rand.Int()), uint(rand.Int())}
	randImageableType := []string{"Bio", "User"}
	randImageName := []string{"img1.jpg", "img2.jpg"}

	oldImage := mockAndInsertImage(db, width, height, randImageableId, randImageableType, randImageName, 2)
	defer destructCreatedObjects(db, oldImage)

	newImage := &Image{
		Id:            oldImage[0].Id,
		Name:          oldImage[1].Name,
		Path:          oldImage[1].Path,
		ImageableId:   oldImage[1].ImageableId,
		ImageableType: oldImage[1].ImageableType,
	}
	defer removeCreatedImageFile([]Image{oldImage[0], *newImage})

	wrongImage := &Image{
		Name: "",
		Path: "",
	}

	_, err = sv.UpdateImage(newImage.Id, &wrongImage.Name, &wrongImage.Path)
	assert.Error(t, err, "Image service update functionality failed")
	assert.ErrorIs(t, err, NameNotFound, "Image service update functionality failed")

	updatedImage, err := sv.UpdateImage(newImage.Id, &newImage.Name, &newImage.Path)

	assert.NoError(t, err, "Image service update user failed")
	assert.Equal(t, newImage.Id, updatedImage.Id, "Image service update bio failed")
	assert.Equal(t, newImage.Name, updatedImage.Name, "Image service update bio failed")

}

func createImageService(db *gorm.DB) ImageServiceInterface {
	return NewImageService(NewImageRepository(db))
}
