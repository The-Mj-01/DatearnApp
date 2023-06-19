package bio

type BioService struct {
	repo BioRepositoryInterface
}

func NewService(repo BioRepositoryInterface) BioServiceInterface {
	return &BioService{
		repo: repo,
	}
}

func (b BioService) GetBioById(id uint) (*Bio, error) {
	return b.repo.GetBioById(id)
}

func (b BioService) GetBioByCountry(countryId uint) (*[]Bio, error) {
	return b.repo.GetBatchesBioByCountry(countryId)
}

func (b BioService) GetBioByCity(cityId uint) (*[]Bio, error) {
	return b.repo.GetBatchesBioByCity(cityId)
}

func (b BioService) GetBioBySex(sexId uint) (*[]Bio, error) {
	return b.repo.GetBatchesBioBySex(sexId)
}

func (b BioService) GetBioByBorn(bornDate int64) (*[]Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioService) GetBioByBornAfter(bornDate int64) (*[]Bio, error) {
	return b.repo.GetBatchesBioByBornAfter(bornDate)
}

func (b BioService) GetBioByCountryCitySex(countryId, cityId, sexId uint) (*[]Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioService) GetBioByCountryCitySexBornAfterDate(countryId, cityId, sexId uint, bornDate int64) (*[]Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioService) CreateBio(description string, socialMedia, country, city, sex uint, born int64) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioService) UpdateBio(description string, socialMedia, country, city, sex uint, born int64) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}
