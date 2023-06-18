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

func (b BioRepository) GetBatchesBioByCountry(countryId uint) (*Bio, error) {
	var bio Bio
	result := b.db.Where("country = ?", countryId).Find(&bio)
	return &bio, result.Error
}

func (b BioRepository) GetBatchesBioByCity(cityId uint) (*Bio, error) {
	var bio Bio
	result := b.db.Where("city = ?", cityId).Find(&bio)
	return &bio, result.Error
}

func (b BioRepository) GetBatchesBioBySex(sexId uint) (*Bio, error) {
	var bio Bio
	result := b.db.Where("sex = ?", sexId).Find(&bio)
	return &bio, result.Error
}

func (b BioRepository) GetBatchesBioByBorn(bornDate time.Time) (*Bio, error) {
	//TODO implement me
	panic("implement me")
}

func (b BioRepository) GetBatchesBioByBornAfter(bornDate time.Time) (*Bio, error) {
	var bio Bio
	result := b.db.Where("born >= ?", bornDate).Find(&bio)
	return &bio, result.Error
}

func (b BioRepository) GetBatchesBioByCountryCitySex(countryId, cityId, sexId uint) (*Bio, error) {
	var bio Bio
	result := b.db.Where("country = ? ", countryId).Where("city", cityId).Where("sex", sexId).Find(&bio)
	return &bio, result.Error
}

func (b BioRepository) GetBatchesBioByCountryCitySexBornAfterDate(countryId, cityId, sexId uint, bornDate time.Time) (*Bio, error) {
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
