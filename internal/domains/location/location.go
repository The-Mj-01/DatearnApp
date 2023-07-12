package location

import "context"

type CountryRepositoryInterface interface {
	GetAllCountries(name *string, limit *int, offset int) *[]Country
	CountryExists(countryId uint) bool
}
type CountryServiceInterface interface {
	GetAllCountries(name *string, limit *int, offset int) (*[]Country, error)
}

type CountryUseCaseInterface interface {
	GetAllCountries(ctx context.Context, token string, request *CountryGetrequest) (*[]Country, error)
}

type CityRepositoryInterface interface {
	GetAllCities(name *string, limit *int, offset int) *[]City
	CityExists(cityId uint) bool
}

type CityServiceInterface interface {
	GetAllCities(name *string, limit *int, offset int) (*[]City, error)
}

type CityUseCaseInterface interface {
	GetAllCities(ctx context.Context, token string, request *CityGetrequest) (*[]City, error)
}
