package location

import (
	"Datearn/pkg/advancedError"
	"context"
)

type CountryUseCase struct {
	sv        CountryServiceInterface
	decoderFn func(ctx context.Context, token string) (uint, error)
}

type CityUseCase struct {
	sv        CityServiceInterface
	decoderFn func(ctx context.Context, token string) (uint, error)
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
	_, err := c.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return c.sv.GetAllCities(request.Name, request.Limit, request.Offset)
}
