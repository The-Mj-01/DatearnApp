package image

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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

	mockedGetImageRequest := mockGetImageRequest(&mockedImage[0].Id, &mockedImage[0].ImageableId, &mockedImage[0].Name, &mockedImage[0].ImageableType, nil, 0)

	fetchedImage, err := useCase.GetAllImage(ctx, "", mockedGetImageRequest)
	assert.NotNil(t, fetchedImage)
	assertImage(t, mockedImage, *fetchedImage)
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

func mockWriteImageRequest(name, imageableType string, imageableId uint) *ImageCreateRequest {
	return &ImageCreateRequest{
		Name:          name,
		ImageableId:   imageableId,
		ImageableType: imageableType,
	}
}

func mockEditImageRequest(id, imageableId *uint, name, imageableType *string) *ImageUpdateRequest {
	return &ImageUpdateRequest{
		Id:            id,
		Name:          name,
		ImageableId:   imageableId,
		ImageableType: imageableType,
	}
}

func mockDeleteImageRequest(id uint) *ImageDeleteRequest {
	return &ImageDeleteRequest{
		Id: id,
	}
}
