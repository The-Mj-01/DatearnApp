package image

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"testing"
)

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Image{})
	return db, err
}

// createImageRepo for testing purpose and return it
func createImageRepo(db *gorm.DB) ImageRepositoryInterface {
	return NewImageRepository(db)
}

// mockAndInsertImage in database for testing purpose
func mockAndInsertImage(db *gorm.DB, width, height []int, count int, imageableId []uint, imageableType []string) []Image {
	img := make([]Image, 0, count)
	i := 0
	for {
		tmpImage := mockImage(width[i], height[i], imageableId[i], imageableType[i])

		res := db.Create(tmpImage)
		if res.Error != nil {
			continue
		}

		img = append(img, *tmpImage)
		i += 1

		if i == count {
			break
		}
	}
	return img
}

// mockImage object and return it
func mockImage(width, height int, imageableId uint, imageableType string) *Image {
	filename, path := mockCreateImageFile(width, height)
	return &Image{
		Name:          filename,
		Path:          path,
		ImageableId:   imageableId,
		ImageableType: imageableType,
	}
}

func mockCreateImageFile(width, height int) (string, string) {
	// Create a new image with the given width and height.
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Define the colors to use.
	blue := color.RGBA{0, 0, 255, 255}

	// Set the color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, blue)
		}
	}

	// Create a file to save the image to.
	file, err := os.Create("img.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode the image as jpeg and write it to the file.
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		panic(err)
	}

	return file.Name(), filepath.Join(".", file.Name())
}

// assertImage check whether they are equal or not
func assertImage(t *testing.T, createdImage, fetchedImage []Image) {
	for index := range createdImage {
		assert.Equal(t, createdImage[index].Id, fetchedImage[index].Id, "Image Repository test: Ids are not equal")
		assert.Equal(t, createdImage[index].Name, fetchedImage[index].Name, "Image Repository test: titles are not equal")

	}
}

// destructCreatedObjects that are created for testing purpose
func destructCreatedObjects[T Image](db *gorm.DB, records []T) {
	for _, record := range records {
		db.Unscoped().Delete(record)
	}
}
