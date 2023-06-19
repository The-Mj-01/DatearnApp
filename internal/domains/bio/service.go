package bio

type BioService struct {
	repo BioRepositoryInterface
}

func NewService(repo BioRepositoryInterface) BioServiceInterface {
	return &BioService{
		repo: repo,
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
	countryExists := b.repo.CountryExists(country)
	if !countryExists {
		return nil, CountryNotFound
	}
	cityExists := b.repo.CityExists(city)
	if !cityExists {
		return nil, CityNotFound
	}
	sexExists := b.repo.SexExists(sex)
	if !sexExists {
		return nil, SexNotFound
	}

	if userId == 0 {
		return nil, UserIdNotFound
	}

	if born == 0 {
		return nil, BornIdNotFound
	}

	tmpBio := &Bio{

		UserId:      userId,
		Description: description,
		Country:     country,
		City:        city,
		Sex:         sex,
		Born:        born,
	}
	bio, err := b.repo.CreateBio(tmpBio)
	return bio, err
}

func (b BioService) UpdateBio(description string, country, city, sex uint, born int64) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}
