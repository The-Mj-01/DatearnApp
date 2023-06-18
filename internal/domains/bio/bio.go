package bio

import (
	"context"
)

type BioRepositoryInterface interface {
	GetBioById(id uint) (*Bio, error)
	GetBatchesBioByCountry(countryId uint) (*[]Bio, error)
	GetBatchesBioByCity(cityId uint) (*[]Bio, error)
	GetBatchesBioBySex(sexId uint) (*[]Bio, error)
	GetBatchesBioByBorn(bornDate int64) (*[]Bio, error)

	//
	GetBatchesBioByBornAfter(bornDate int64) (*[]Bio, error)
	GetBatchesBioByCountryCitySex(countryId, cityId, sexId uint) (*[]Bio, error)
	GetBatchesBioByCountryCitySexBornAfterDate(countryId, cityId, sexId uint, bornDate int64) (*[]Bio, error)
	CreateBio(bio *Bio) (*Bio, error)
	UpdateBio(oldBio, newBio *Bio) (*Bio, error)
}

type BioServiceInterface interface {
	GetBioById(id uint) (*Bio, error)
	GetBioByCountry(countryId uint) (*[]Bio, error)
	GetBioByCity(cityId uint) (*[]Bio, error)
	GetBioBySex(sexId uint) (*[]Bio, error)
	GetBioByBorn(bornDate int64) (*[]Bio, error)

	//
	GetBioByBornAfter(bornDate int64) (*[]Bio, error)
	GetBioByCountryCitySex(countryId, cityId, sexId uint) (*[]Bio, error)
	GetBioByCountryCitySexBornAfterDate(countryId, cityId, sexId uint, bornDate int64) (*[]Bio, error)
	CreateBio(description string, socialMedia, country, city, sex uint, born int64) (*Bio, error)
	UpdateBio(description string, socialMedia, country, city, sex uint, born int64) (*Bio, error)
}

type BioUseCaseInterface interface {
	WriteBio(ctx context.Context, request *BioCreateRequest) (*Bio, error)
	GetBio(ctx context.Context, request *BioGetSingleRequest) (*Bio, error)
	UpdateBio(ctx context.Context, request *BioUpdateRequest) (*Bio, error)
}
