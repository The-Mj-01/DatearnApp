package image

import "context"

type ImageRepositoryInterface interface {
	GetAllImage(id, imageableId *uint, name, imageableType *string, limit *int, offset int) *[]Image
	CreateImage(imageableId uint, name, path, imageableType string) (*Image, error)
	UpdateImage(oldImage, newImage *Image) (*Image, error)
	DeleteImage(image *Image) (*Image, error)
}

type ImageServiceInterface interface {
	GetAllImage(id, imageableId *uint, name, imageableType *string, limit *int, offset int) (*[]Image, error)
	CreateImage(imageableId uint, name, path, imageableType string) (*Image, error)
	UpdateImage(id uint, name, path *string) (*Image, error)
	DeleteImage(imageId *uint) (*Image, error)
}

type ImageUseCaseInterface interface {
	GetAllImage(ctx context.Context, token string, request *ImageGetRequest) (*[]Image, error)
	CreateImage(ctx context.Context, token string, request *ImageCreateRequest) (*Image, error)
	UpdateImage(ctx context.Context, token string, request *ImageUpdateRequest) (*Image, error)
	DeleteImage(ctx context.Context, token string, request *ImageDeleteRequest) (*Image, error)
}
