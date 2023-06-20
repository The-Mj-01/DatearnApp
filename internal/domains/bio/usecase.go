package bio

import "context"

// BioUseCase is a struct which satisfies user use case interface functionalities
type BioUseCase struct {
	sv BioServiceInterface
}

// NewUserUseCase and return it
func NewUserUseCase(sv BioServiceInterface) BioUseCaseInterface {
	return &BioUseCase{
		sv: sv,
	}
}

func (b BioUseCase) WriteBio(ctx context.Context, token string, request *BioCreateRequest) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioUseCase) GetBio(ctx context.Context, token string, request *BioGetSingleRequest) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioUseCase) UpdateBio(ctx context.Context, token string, request *BioUpdateRequest) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}
