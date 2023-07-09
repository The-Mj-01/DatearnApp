package image

import (
	"Datearn/pkg/advancedError"
	"Datearn/pkg/userHandler"
	"context"
)

// ImageUseCase is a struct which satisfies user use case interface functionalities
type ImageUseCase struct {
	sv        ImageServiceInterface
	decoderFn func(ctx context.Context, token string) (uint, error)
}

// NewImageUseCase and return it
func NewImageUseCase(sv ImageServiceInterface, decoderFn func(ctx context.Context, token string) (uint, error)) ImageUseCaseInterface {
	if decoderFn == nil {
		decoderFn = userHandler.ExtractUserIdFromToken
	}

	return &ImageUseCase{
		sv:        sv,
		decoderFn: decoderFn,
	}
}

func (i *ImageUseCase) GetAllImage(ctx context.Context, token string, request *ImageGetRequest) (*[]Image, error) {
	_, err := i.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return i.sv.GetAllImage(request.Id, request.ImageableId, request.Name, request.ImageableType, request.Limit, request.Offset)
}

func (i *ImageUseCase) CreateImage(ctx context.Context, token string, request *ImageCreateRequest) (*Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i *ImageUseCase) UpdateImage(ctx context.Context, token string, request *ImageUpdateRequest) (*Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i *ImageUseCase) DeleteImage(ctx context.Context, token string, request *ImageDeleteRequest) (*Image, error) {
	//TODO implement me
	panic("implement me")
}
