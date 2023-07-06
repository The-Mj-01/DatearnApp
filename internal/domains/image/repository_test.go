package image

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const imagePath = "../../../static/images/"

func TestImageRepository_GetAllImage(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createImageRepo(db)

	width := []int{200, 200, 200, 200, 200}
	height := []int{200, 200, 200, 200, 200}
	randImageableId := []uint{uint(rand.Int()), uint(rand.Int()), uint(rand.Int()), uint(rand.Int()), uint(rand.Int())}
	randImageableType := []string{"Bio", "Bio", "Bio", "Bio", "Bio"}
	randImageName := []string{"img1.jpg", "img2.jpg", "img3.jpg", "img4.jpg", "img5.jpg"}

	img := mockAndInsertImage(db, width, height, randImageableId, randImageableType, randImageName, 5)
	defer destructCreatedObjects(db, img)
	defer removeCreatedImageFile(img)

	fetchedImage := repo.GetAllImage(nil, nil, nil, nil, nil, 10)
	assert.Equal(t, len(*fetchedImage), 0, "Fetched image  are not equal")

	limit := 1
	fetchedImage = repo.GetAllImage(nil, nil, nil, nil, &limit, 0)
	assert.Equal(t, len(*fetchedImage), limit, "one image must be fetched")

	falseType := "WrongType"
	fetchedImage = repo.GetAllImage(nil, nil, nil, &falseType, nil, 0)
	assert.Equal(t, len(*fetchedImage), 0, "zero image must be fetched")

	fetchedImage = repo.GetAllImage(nil, nil, &img[0].Name, nil, nil, 0)
	assert.NotZero(t, len(*fetchedImage), "Zero image  fetched")
	assertImage(t, img, *fetchedImage)

	fetchedImage = repo.GetAllImage(nil, nil, nil, nil, nil, 0)
	assert.NotZero(t, len(*fetchedImage), "Zero image  fetched")
	assert.Equal(t, len(*fetchedImage), 5, "Fetched img  are not equal")
	assertImage(t, img, *fetchedImage)

}

func TestImageRepository_CreateImage(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createImageRepo(db)

	width := 200
	height := 200
	randImageableId := uint(rand.Int())
	randImageableType := "Bio"
	randImageName := "img1.jpg"

	img := mockImage(width, height, randImageableId, randImageableType, randImageName)

	createdImage, err := repo.CreateImage(img.ImageableId, img.Name, img.Path, img.ImageableType)
	defer destructCreatedObjects(db, []Image{*createdImage})
	defer removeCreatedImageFile([]Image{*createdImage})

	assert.NoError(t, err, "Image creation in repository failed")
	assert.Equal(t, img.Name, createdImage.Name, "Image Repository test: titles are not equal")
}

func TestImageRepository_UpdateImage(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createImageRepo(db)

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

	_, err = repo.UpdateImage(&oldImage[0], newImage)
	assert.NoError(t, err, "Image Update operation failed")

	fetchImage := new(Image)
	db.Where("id = ?", oldImage[0].Id).First(fetchImage)

	assert.Equal(t, newImage.Id, fetchImage.Id, "Image Update operation failed")
	assert.Equal(t, newImage.Name, fetchImage.Name, "Image Update operation failed")
	assert.Equal(t, newImage.Path, fetchImage.Path, "Image Update operation failed")
	assert.Equal(t, newImage.ImageableId, fetchImage.ImageableId, "Image Update operation failed")
	assert.Equal(t, newImage.ImageableType, fetchImage.ImageableType, "Image Update operation failed")

}

func TestImageRepository_DeleteImage(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createImageRepo(db)

	width := []int{200, 200, 200, 200, 200}
	height := []int{200, 200, 200, 200, 200}
	randImageableId := []uint{uint(rand.Int()), uint(rand.Int()), uint(rand.Int()), uint(rand.Int()), uint(rand.Int())}
	randImageableType := []string{"Bio", "Bio", "Bio", "Bio", "Bio"}
	randImageName := []string{"img1.jpg", "img2.jpg", "img3.jpg", "img4.jpg", "img5.jpg"}

	img := mockAndInsertImage(db, width, height, randImageableId, randImageableType, randImageName, 5)
	defer destructCreatedObjects(db, img)
	defer removeCreatedImageFile(img)

	deletedImage, err := repo.DeleteImage(&img[0])
	assertImage(t, []Image{*deletedImage}, []Image{img[0]})
	fetchUser := new(Image)
	result := db.Where("id = ?", img[0].Id).First(fetchUser)

	assert.Error(t, result.Error, "Image Delete operation failed")
}

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
func mockAndInsertImage(db *gorm.DB, width, height []int, imageableId []uint, imageableType, name []string, count int) []Image {
	img := make([]Image, 0, count)
	i := 0
	for {
		tmpImage := mockImage(width[i], height[i], imageableId[i], imageableType[i], name[i])

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
func mockImage(width, height int, imageableId uint, imageableType, name string) *Image {
	filename, path := mockCreateImageFile(width, height, name)
	return &Image{
		Name:          filename,
		Path:          path,
		ImageableId:   imageableId,
		ImageableType: imageableType,
	}
}

func mockCreateImageFile(width, height int, name string) (string, string) {
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

	pathDir, err := createDirectoryFromDate()

	if err != nil {
		panic(err)
	}

	// Create a file to save the image to.
	file, err := os.Create(imagePath + pathDir + name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode the image as jpeg and write it to the file.
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		panic(err)
	}

	return file.Name(), filepath.Join("../../static/images/", file.Name())
}

func createDirectoryFromDate() (string, error) {
	// Get the current date.
	t := time.Now()

	// Format the date as year/month/day.
	date := t.Format("2006/01/02")

	// Print the date.
	//fmt.Println(date)

	// Create the directories with the date as the path.
	err := os.MkdirAll(imagePath+date, 0755)
	if err != nil {
		// Print the error if any.
		fmt.Println(err)
		return "", err
	}

	// Print a success message.
	//fmt.Println("Directories have been created successfully!")
	return date + "/", nil
}

// assertImage check whether they are equal or not
func assertImage(t *testing.T, createdImage, fetchedImage []Image) {
	for index := range fetchedImage {
		assert.Equal(t, createdImage[index].Id, fetchedImage[index].Id, "Image Repository test: Ids are not equal")
		assert.Equal(t, createdImage[index].Name, fetchedImage[index].Name, "Image Repository test: titles are not equal")

	}
}

// destructCreatedObjects that are created for testing purpose
func destructCreatedObjects(db *gorm.DB, records []Image) {
	for _, record := range records {
		db.Unscoped().Delete(record)
	}
}

func removeCreatedImageFile(imgs []Image) {
	for _, img := range imgs {
		err := os.Remove(img.Path)
		if err != nil {
			log.Fatal(err)
		}
	}
}
