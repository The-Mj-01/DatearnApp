package image

import (
	"context"
	"gorm.io/gorm"
)

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
