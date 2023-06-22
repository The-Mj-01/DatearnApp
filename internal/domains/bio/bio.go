package bio

import (
	"context"
)

type BioRepositoryInterface interface {
	GetBioByUserId(userId uint) (*Bio, error)
	GetBioById(id uint) (*Bio, error)
	GetBatchesBioByCountry(countryId uint) (*[]Bio, error)
	GetBatchesBioByCity(cityId uint) (*[]Bio, error)
	GetBatchesBioBySex(sexId uint) (*[]Bio, error)
	GetBatchesBioByBorn(bornDate int64) (*[]Bio, error)
	CountryExists(countryId uint) bool
	CityExists(cityId uint) bool
	SexExists(sexId uint) bool
	//
	GetBatchesBioByBornAfter(bornDate int64) (*[]Bio, error)
	GetBatchesBioByCountryCitySex(countryId, cityId, sexId uint) (*[]Bio, error)
	GetBatchesBioByCountryCitySexBornAfterDate(countryId, cityId, sexId uint, bornDate int64) (*[]Bio, error)
	CreateBio(bio *Bio) (*Bio, error)
	UpdateBio(oldBio, newBio *Bio) (*Bio, error)
}

type BioServiceInterface interface {
	GetBioByUserId(userId uint) (*Bio, error)
	GetBioById(id uint) (*Bio, error)
	GetBioByCountry(countryId uint) (*[]Bio, error)
	GetBioByCity(cityId uint) (*[]Bio, error)
	GetBioBySex(sexId uint) (*[]Bio, error)
	GetBioByBorn(bornDate int64) (*[]Bio, error)

	//
	GetBioByBornAfter(bornDate int64) (*[]Bio, error)
	GetBioByCountryCitySex(countryId, cityId, sexId uint) (*[]Bio, error)
	GetBioByCountryCitySexBornAfterDate(countryId, cityId, sexId uint, bornDate int64) (*[]Bio, error)
	CreateBio(description string, userId, country, city, sex uint, born int64) (*Bio, error)
	UpdateBio(userId uint, description string, country, city, sex uint, born int64) (*Bio, error)
}

type BioUseCaseInterface interface {
	WriteBio(ctx context.Context, token string, request *BioCreateRequest) (*Bio, error)
	GetBio(ctx context.Context, token string, request *BioGetSingleRequest) (*Bio, error)
	UpdateBio(ctx context.Context, token string, request *BioUpdateRequest) (*Bio, error)
}

type CountryRepositoryInterface interface {
	GetAllCountries(name *string, limit *int, offset int) (*[]Country, error)
}

type CountryServiceInterface interface {
	GetAllCountries(name *string, limit *int, offset int) (*[]Country, error)
}

type CountryUseCaseInterface interface {
	GetAllCountries(ctx context.Context, token string, request *CountryGetrequest) (*[]Country, error)
}

type CityRepositoryInterface interface {
	GetAllCities(name *string, limit *int, offset int) (*[]City, error)
}

type CityServiceInterface interface {
	GetAllCities(name *string, limit *int, offset int) (*[]City, error)
}

type CityUseCaseInterface interface {
	GetAllCities(ctx context.Context, token string, request *CityGetrequest) (*[]City, error)
}
