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

type CountryUseCase struct {
	sv        CountryServiceInterface
	decoderFn func(ctx context.Context, token string) (uint, error)
}

type CityUseCase struct {
	sv        CityServiceInterface
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
	userId, err := b.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return b.sv.GetBioByUserId(userId)
}

func (b BioUseCase) UpdateBio(ctx context.Context, token string, request *BioUpdateRequest) (*Bio, error) {
	userId, err := b.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}

	return b.sv.UpdateBio(userId, request.Description, request.Country, request.City, request.Sex, request.Born)
}

// NewCountryUseCase and return it
func NewCountryUseCase(sv CountryServiceInterface, decoderFn func(ctx context.Context, token string) (uint, error)) CountryUseCaseInterface {
	return &CountryUseCase{
		sv:        sv,
		decoderFn: decoderFn,
	}
}

// GetAllCountries and return them
func (c CountryUseCase) GetAllCountries(ctx context.Context, token string, request *CountryGetrequest) (*[]Country, error) {
	_, err := c.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return c.sv.GetAllCountries(request.Name, request.Limit, request.Offset)
}

// NewCityUseCase and return it
func NewCityUseCase(sv CityServiceInterface, decoderFn func(ctx context.Context, token string) (uint, error)) CityUseCaseInterface {
	return &CityUseCase{
		sv:        sv,
		decoderFn: decoderFn,
	}
}

// GetAllCities and return them
func (c *CityUseCase) GetAllCities(ctx context.Context, token string, request *CityGetrequest) (*[]City, error) {
	//TODO implement me
	panic("implement me")
}
