package bio

import (
	"gorm.io/gorm"
	"time"
)

// BioRepository which implements repository interface
type BioRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) BioRepositoryInterface {
	return &BioRepository{
		db: db,
	}
}

func (b BioRepository) GetBioById(id uint) (*Bio, error) {
	var bio Bio
	result := b.db.Where("id = ?", id).First(&bio)
	return &bio, result.Error
}

func (b BioRepository) GetBioByCountry(countryId uint) (*Bio, error) {
	var bio Bio
	result := b.db.Where("country = ?", countryId).First(&bio)
	return &bio, result.Error
}

func (b BioRepository) GetBioByCity(cityId uint) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioRepository) GetBioBySex(sexId uint) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioRepository) GetBioByBorn(bornDate time.Time) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioRepository) GetBioByBornAfter(bornDate time.Time) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioRepository) GetBioByCountryCitySex(countryId, cityId, sexId uint) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioRepository) GetBioByCountryCitySexBornAfterDate(countryId, cityId, sexId uint, bornDate time.Time) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioRepository) CreateBio(bio *Bio) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioRepository) UpdateBio(bio *Bio) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}
