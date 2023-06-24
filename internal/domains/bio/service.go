package bio

import (
	"Datearn/pkg/authorization"
	"Datearn/pkg/bioHandler"
)

// userAddrAuthorizerField is the field which authorization should get done and checked with that
const userAddrAuthorizerField string = "UserId"

type BioService struct {
	repo             BioRepositoryInterface
	countryValidator func(id uint) error
	cityValidator    func(id uint) error
}

func NewBioService(repo BioRepositoryInterface, countryValidator, cityValidator func(id uint) error) BioServiceInterface {

	if countryValidator == nil {
		countryValidator = bioHandler.CountryIsValid
	}

	if cityValidator == nil {
		cityValidator = bioHandler.CityIsValid
	}

	return &BioService{
		repo:             repo,
		countryValidator: countryValidator,
		cityValidator:    cityValidator,
	}
}

func (b BioService) GetBioByUserId(userId uint) (*Bio, error) {
	return b.repo.GetBioByUserId(userId)
}

func (b BioService) GetBioById(id uint) (*Bio, error) {
	return b.repo.GetBioById(id)
}

func (b BioService) GetBioByCountry(countryId uint) (*[]Bio, error) {
	bios, err := b.repo.GetBatchesBioByCountry(countryId)
	if len(*bios) == 0 {
		return nil, BioDoesntExists
	}
	return bios, err
}

func (b BioService) GetBioByCity(cityId uint) (*[]Bio, error) {
	bios, err := b.repo.GetBatchesBioByCity(cityId)
	if len(*bios) == 0 {
		return nil, BioDoesntExists
	}
	return bios, err
}

func (b BioService) GetBioBySex(sexId uint) (*[]Bio, error) {
	bios, err := b.repo.GetBatchesBioBySex(sexId)
	if len(*bios) == 0 {
		return nil, BioDoesntExists
	}
	return bios, err
}

func (b BioService) GetBioByBorn(bornDate int64) (*[]Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioService) GetBioByBornAfter(bornDate int64) (*[]Bio, error) {
	bios, err := b.repo.GetBatchesBioByBornAfter(bornDate)
	if len(*bios) == 0 {
		return nil, BioDoesntExists
	}
	return bios, err
}

func (b BioService) GetBioByCountryCitySex(countryId, cityId, sexId uint) (*[]Bio, error) {
	bios, err := b.repo.GetBatchesBioByCountryCitySex(countryId, cityId, sexId)
	if len(*bios) == 0 {
		return nil, BioDoesntExists
	}
	return bios, err
}

func (b BioService) GetBioByCountryCitySexBornAfterDate(countryId, cityId, sexId uint, bornDate int64) (*[]Bio, error) {
	bios, err := b.repo.GetBatchesBioByCountryCitySexBornAfterDate(countryId, cityId, sexId, bornDate)
	if len(*bios) == 0 {
		return nil, BioDoesntExists
	}
	return bios, err
}

func (b BioService) CreateBio(description string, userId, country, city, sex uint, born int64) (*Bio, error) {
	if description == "" {
		return nil, DescripitonNotFound
	}

	tmpBio := &Bio{

		UserId:      userId,
		Description: description,
		Country:     country,
		City:        city,
		Sex:         sex,
		Born:        born,
	}

	err := b.bioCanCreated(tmpBio)
	if err != nil {
		return nil, err
	}

	sexExists := b.repo.SexExists(sex)
	if !sexExists {
		return nil, SexNotFound
	}

	if userId == 0 {
		return nil, UserIdNotFound
	}

	if born == 0 {
		return nil, BornNotFound
	}

	bio, err := b.repo.CreateBio(tmpBio)
	return bio, err
}

func (b BioService) UpdateBio(userId uint, description string, country, city, sex uint, born int64) (*Bio, error) {
	if description == "" {
		return nil, DescripitonNotFound
	}

	newBio := &Bio{

		UserId:      userId,
		Description: description,
		Country:     country,
		City:        city,
		Sex:         sex,
		Born:        born,
	}

	err := b.bioCanCreated(newBio)
	if err != nil {
		return nil, err
	}

	sexExists := b.repo.SexExists(sex)
	if !sexExists {
		return nil, SexNotFound
	}

	if userId == 0 {
		return nil, UserIdNotFound
	}

	if born == 0 {
		return nil, BornNotFound
	}

	bio, err := b.GetBioByUserId(userId)

	if err = authorization.SimpleFieldAuthorization(*bio, userAddrAuthorizerField, userId, YouAreNotAllowed); err != nil {
		return nil, err
	}

	return b.repo.UpdateBio(bio, newBio)
}

// bioCanCreated checks whether payment is allowed to insert in db or not!
func (b *BioService) bioCanCreated(bio *Bio) error {
	if err := b.countryValidator(bio.Country); err != nil {
		return err
	}

	if err := b.cityValidator(bio.City); err != nil {
		return err
	}

	return nil
}
