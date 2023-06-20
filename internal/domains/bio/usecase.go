package bio

import (
	"Datearn/pkg/advancedError"
	"Datearn/pkg/userHandler"
	"context"
)

// BioUseCase is a struct which satisfies user use case interface functionalities
type BioUseCase struct {
	sv        BioServiceInterface
	decoderFn func(ctx context.Context, token string) (uint, error)
}

// NewBioUseCase and return it
func NewBioUseCase(sv BioServiceInterface, decoderFn func(ctx context.Context, token string) (uint, error)) BioUseCaseInterface {
	if decoderFn == nil {
		decoderFn = userHandler.ExtractUserIdFromToken
	}

	return &BioUseCase{
		sv:        sv,
		decoderFn: decoderFn,
	}
}

func (b BioUseCase) WriteBio(ctx context.Context, token string, request *BioCreateRequest) (*Bio, error) {
	userId, err := b.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return b.sv.CreateBio(request.Description, userId, request.Country, request.City, request.Sex, request.Born)
}

func (b BioUseCase) GetBio(ctx context.Context, token string, request *BioGetSingleRequest) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioUseCase) UpdateBio(ctx context.Context, token string, request *BioUpdateRequest) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}
