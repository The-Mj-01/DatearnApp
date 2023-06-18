package bio

import "time"

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

func (b BioService) GetBioByCountry(countryId uint) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioService) GetBioByCity(cityId uint) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioService) GetBioBySex(sexId uint) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioService) GetBioByBorn(bornDate time.Time) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioService) GetBioByBornAfter(bornDate time.Time) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioService) GetBioByCountryCitySex(countryId, cityId, sexId uint) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioService) GetBioByCountryCitySexBornAfterDate(countryId, cityId, sexId uint, bornDate time.Time) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioService) CreateBio(description string, socialMedia, country, city, sex uint, born time.Time) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioService) UpdateBio(description string, socialMedia, country, city, sex uint, born time.Time) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}
