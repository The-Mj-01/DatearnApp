package bio

import (
	"context"
	"time"
)

type BioRepositoryInterface interface {
	GetBioById(id uint) (*Bio, error)
	GetBatchesBioByCountry(countryId uint) (*[]Bio, error)
	GetBatchesBioByCity(cityId uint) (*[]Bio, error)
	GetBatchesBioBySex(sexId uint) (*[]Bio, error)
	GetBatchesBioByBorn(bornDate time.Time) (*[]Bio, error)

	//
	GetBatchesBioByBornAfter(bornDate time.Time) (*[]Bio, error)
	GetBatchesBioByCountryCitySex(countryId, cityId, sexId uint) (*[]Bio, error)
	GetBatchesBioByCountryCitySexBornAfterDate(countryId, cityId, sexId uint, bornDate time.Time) (*[]Bio, error)
	CreateBio(bio *Bio) (*Bio, error)
	UpdateBio(bio *Bio) (*Bio, error)
}

type BioServiceInterface interface {
	GetBioById(id uint) (*Bio, error)
	GetBioByCountry(countryId uint) (*Bio, error)
	GetBioByCity(cityId uint) (*Bio, error)
	GetBioBySex(sexId uint) (*Bio, error)
	GetBioByBorn(bornDate time.Time) (*Bio, error)

	//
	GetBioByBornAfter(bornDate time.Time) (*Bio, error)
	GetBioByCountryCitySex(countryId, cityId, sexId uint) (*Bio, error)
	GetBioByCountryCitySexBornAfterDate(countryId, cityId, sexId uint, bornDate time.Time) (*Bio, error)
	CreateBio(description string, socialMedia, country, city, sex uint, born time.Time) (*Bio, error)
	UpdateBio(description string, socialMedia, country, city, sex uint, born time.Time) (*Bio, error)
}

type BioUseCaseInterface interface {
	WriteBio(ctx context.Context, request *BioCreateRequest) (*Bio, error)
	GetBio(ctx context.Context, request *BioGetSingleRequest) (*Bio, error)
	UpdateBio(ctx context.Context, request *BioUpdateRequest) (*Bio, error)
}
