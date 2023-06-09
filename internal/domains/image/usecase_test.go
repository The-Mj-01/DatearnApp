package image

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"image"
	"math/rand"
	"testing"
)

func TestImageUseCase_GetAllImage(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createImageUseCase(db, randUserId)

	width := []int{200, 200, 200, 200, 200}
	height := []int{200, 200, 200, 200, 200}
	randImageableId := []uint{uint(rand.Int()), uint(rand.Int()), uint(rand.Int()), uint(rand.Int()), uint(rand.Int())}
	randImageableType := []string{"Bio", "Bio", "Bio", "Bio", "Bio"}
	randImageName := []string{"img1.jpg", "img2.jpg", "img3.jpg", "img4.jpg", "img5.jpg"}

	mockedImage := mockAndInsertImage(db, width, height, randImageableId, randImageableType, randImageName, 5)
	assert.Equal(t, len(mockedImage), 5, "Mocking products failed")
	defer destructCreatedObjects(db, mockedImage)
	defer removeCreatedImageFile(mockedImage)

	mockedGetImageRequest := mockGetImageRequest(&mockedImage[0].Id, &mockedImage[0].ImageableId, &mockedImage[0].Name, &mockedImage[0].ImageableType, nil, 0)

	fetchedImage, err := useCase.GetAllImage(ctx, "", mockedGetImageRequest)
	assert.NotNil(t, fetchedImage)
	assertImage(t, mockedImage, *fetchedImage)
}

func TestImageUseCase_CreateImage(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createImageUseCase(db, randUserId)

	width := 200
	height := 200
	randImageableId := uint(rand.Int())
	randImageableType := "Bio"
	randImageName := "img1.jpg"

	img := createImage(width, height)
	mockedRequest := mockWriteImageRequest(randImageName, randImageableType, randImageableId, img)
	result, err := useCase.CreateImage(ctx, "", mockedRequest)
	defer destructCreatedObjects(db, []Image{*result})
	defer removeCreatedImageFile([]Image{*result})

	assert.NoError(t, err, "Image creation failed in address use-case")
	assert.Equal(t, result.Name, mockedRequest.Name, "Image creation failed in bio use-case")
}

func TestImageUseCase_UpdateImage(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createImageUseCase(db, randUserId)

	width := []int{200, 400}
	height := []int{200, 400}
	randImageableId := []uint{uint(rand.Int()), uint(rand.Int())}
	randImageableType := []string{"Bio", "User"}
	randImageName := []string{"img1.jpg"}

	oldImage := mockAndInsertImage(db, width, height, randImageableId, randImageableType, randImageName, 1)
	defer destructCreatedObjects(db, oldImage)

	width2 := 200
	height2 := 200
	randImageableId2 := uint(rand.Int())
	randImageableType2 := "User"
	randImageName2 := "img2.jpg"
	newimg := mockImage(width2, height2, randImageableId2, randImageableType2, randImageName2)

	newImage := &Image{
		Id:            oldImage[0].Id,
		Name:          newimg.Name,
		Path:          newimg.Path,
		ImageableId:   newimg.ImageableId,
		ImageableType: newimg.ImageableType,
	}
	defer removeCreatedImageFile([]Image{oldImage[0], *newImage})

	//wrongImage := &Image{
	//	Name: "",
	//	Path: "",
	//}

	img := createImage(width2, height2)
	mockedEditRequest := mockEditImageRequest(&newImage.Id, &newImage.ImageableId, &newImage.Name, &newImage.ImageableType, img)
	editedImage, err := useCase.UpdateImage(ctx, "", mockedEditRequest)

	defer destructCreatedObjects(db, []Image{*editedImage})
	assert.NoError(t, err, "Image use-case update functionality failed")

	assert.Equal(t, *mockedEditRequest.Name, editedImage.Name, "Image use-case update functionality failed")
}

func TestImageUseCase_DeleteImage(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createImageUseCase(db, randUserId)

	width := []int{200}
	height := []int{200}
	randImageableId := []uint{uint(rand.Int())}
	randImageableType := []string{"Bio"}
	randImageName := []string{"img1.jpg"}

	mockedImage := mockAndInsertImage(db, width, height, randImageableId, randImageableType, randImageName, 1)
	defer destructCreatedObjects(db, mockedImage)

	mockedDeleteRequest := mockDeleteImageRequest(mockedImage[0].Id)

	deletedImage, err := useCase.DeleteImage(ctx, "", mockedDeleteRequest)
	assert.NoError(t, err, "Deleting user name failed")

	assertImage(t, mockedImage, []Image{*deletedImage})

}

func createImageUseCase(db *gorm.DB, userId uint) ImageUseCaseInterface {
	return NewImageUseCase(NewImageService(NewImageRepository(db)), func(ctx context.Context, token string) (uint, error) {
		return userId, nil
	})
}

func mockGetImageRequest(id, imageableId *uint, name, imageableType *string, limit *int, offset int) *ImageGetRequest {
	return &ImageGetRequest{
		Id:            id,
		Name:          name,
		ImageableId:   imageableId,
		ImageableType: imageableType,
		Limit:         limit,
		Offset:        offset,
	}
}

func mockWriteImageRequest(name, imageableType string, imageableId uint, img *image.RGBA) *ImageCreateRequest {
	return &ImageCreateRequest{
		Name:          name,
		ImageableId:   imageableId,
		ImageableType: imageableType,
		Img:           img,
	}
}

func mockEditImageRequest(id, imageableId *uint, name, imageableType *string, img *image.RGBA) *ImageUpdateRequest {
	return &ImageUpdateRequest{
		Id:            id,
		Name:          name,
		ImageableId:   imageableId,
		ImageableType: imageableType,
		Img:           img,
	}
}

func mockDeleteImageRequest(id uint) *ImageDeleteRequest {
	return &ImageDeleteRequest{
		Id: id,
	}
}
