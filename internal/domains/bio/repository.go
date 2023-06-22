package bio

import (
	"errors"
	"gorm.io/gorm"
)

// BioRepository which implements repository interface
type BioRepository struct {
	db *gorm.DB
}

// CountryRepository which implements repository interface
type CountryRepository struct {
	db *gorm.DB
}

// CityRepository which implements repository interface
type CityRepository struct {
	db *gorm.DB
}

// NewBioRepository instantiates and returns new repository
func NewBioRepository(db *gorm.DB) BioRepositoryInterface {
	return &BioRepository{
		db: db,
	}
}

// GetBioByUserId and return it
func (b BioRepository) GetBioByUserId(userId uint) (*Bio, error) {
	var bio Bio
	result := b.db.Where(" user_id = ?", userId).First(&bio)
	return &bio, result.Error
}

// GetBioById and return it
func (b BioRepository) GetBioById(id uint) (*Bio, error) {
	var bio Bio
	result := b.db.Where("id = ?", id).First(&bio)
	return &bio, result.Error
}

// GetBatchesBioByCountry and return it
func (b BioRepository) GetBatchesBioByCountry(countryId uint) (*[]Bio, error) {
	var bio []Bio
	result := b.db.Where("country = ?", countryId).Find(&bio)
	return &bio, result.Error
}

// GetBatchesBioByCity and return it
func (b BioRepository) GetBatchesBioByCity(cityId uint) (*[]Bio, error) {
	var bio []Bio
	result := b.db.Where("city = ?", cityId).Find(&bio)
	return &bio, result.Error
}

// GetBatchesBioBySex and return it
func (b BioRepository) GetBatchesBioBySex(sexId uint) (*[]Bio, error) {
	var bio []Bio
	result := b.db.Where("sex = ?", sexId).Find(&bio)
	return &bio, result.Error
}

// GetBatchesBioByBorn and return it
func (b BioRepository) GetBatchesBioByBorn(bornDate int64) (*[]Bio, error) {
	//TODO implement me
	panic("implement me")
}

// CountryExists and return is Exists
func (b BioRepository) CountryExists(countryId uint) bool {
	var country Country
	result := b.db.Where("id = ?", countryId).First(&country)
	return result.Error == nil && !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// CityExists and return is Exists
func (b BioRepository) CityExists(cityId uint) bool {
	var city City
	result := b.db.Where("id = ?", cityId).First(&city)
	return result.Error == nil && !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (b BioRepository) SexExists(sexId uint) bool {
	var sex Sex
	result := b.db.Where("id = ?", sexId).First(&sex)
	return result.Error == nil && !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// GetBatchesBioByBornAfter and return it
func (b BioRepository) GetBatchesBioByBornAfter(bornDate int64) (*[]Bio, error) {
	var bio []Bio
	result := b.db.Where("born >= ?", bornDate).Find(&bio)
	return &bio, result.Error
}

// GetBatchesBioByCountryCitySex and return it
func (b BioRepository) GetBatchesBioByCountryCitySex(countryId, cityId, sexId uint) (*[]Bio, error) {
	var bio []Bio
	result := b.db.Where("country = ? ", countryId).Where("city", cityId).Where("sex", sexId).Find(&bio)
	return &bio, result.Error
}

// GetBatchesBioByCountryCitySexBornAfterDate and return it
func (b BioRepository) GetBatchesBioByCountryCitySexBornAfterDate(countryId, cityId, sexId uint, bornDate int64) (*[]Bio, error) {
	var bio []Bio
	result := b.db.Where("country = ? ", countryId).Where("city", cityId).Where("sex", sexId).Where("born >= ?", bornDate).Find(&bio)
	return &bio, result.Error
}

// CreateBio that doesn't already exist in database
func (b BioRepository) CreateBio(bio *Bio) (*Bio, error) {
	result := b.db.Create(bio)
	return bio, result.Error
}

// UpdateBio that already exists in database
func (b BioRepository) UpdateBio(oldBio, newBio *Bio) (*Bio, error) {
	if newBio.Country != 0 {
		oldBio.Country = newBio.Country
	}
	if newBio.City != 0 {
		oldBio.City = newBio.City
	}
	if newBio.Sex != 0 {
		oldBio.Sex = newBio.Sex
	}
	if newBio.Born != 0 {
		oldBio.Born = newBio.Born
	}

	if newBio.Description != "" {
		oldBio.Description = newBio.Description
	}

	result := b.db.Save(oldBio)

	return oldBio, result.Error
}

// NewCountryRepository instantiates and returns new repository
func NewCountryRepository(db *gorm.DB) CountryRepositoryInterface {
	return &CountryRepository{
		db: db,
	}
}

// GetAllCountry and return it
func (c *CountryRepository) GetAllCountries(name *string, limit *int, offset int) (*[]Country, error) {
	var countries *[]Country
	db := c.db

	if name != nil {
		db = db.Where("name LIKE ?", *name)
	}

	if limit != nil {
		db = db.Limit(*limit)
	}

	db.Offset(offset).Find(&countries)
	return countries, db.Error
}

// NewCityRepository instantiates and returns new repository
func NewCityRepository(db *gorm.DB) CityRepositoryInterface {
	return &CityRepository{
		db: db,
	}
}

// GetAllCity and return it
func (c *CityRepository) GetAllCities(name *string, limit *int, offset int) (*[]City, error) {
	var cities *[]City
	db := c.db

	if name != nil {
		db = db.Where("name Like ?", *name)
	}

	if limit != nil {
		db = db.Limit(*limit)
	}

	db.Offset(offset).Find(&cities)
	return cities, db.Error
}
