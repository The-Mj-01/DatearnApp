package location

import (
	"errors"
	"gorm.io/gorm"
)

// CountryRepository which implements repository interface
type CountryRepository struct {
	db *gorm.DB
}

// CityRepository which implements repository interface
type CityRepository struct {
	db *gorm.DB
}

// NewCountryRepository instantiates and returns new repository
func NewCountryRepository(db *gorm.DB) CountryRepositoryInterface {
	return &CountryRepository{
		db: db,
	}
}

// GetAllCountry and return it
func (c *CountryRepository) GetAllCountries(name *string, limit *int, offset int) *[]Country {
	var countries *[]Country
	db := c.db

	if name != nil {
		db = db.Where("name LIKE ?", *name)
	}

	if limit != nil {
		db = db.Limit(*limit)
	}

	db.Offset(offset).Find(&countries)
	return countries
}

// NewCityRepository instantiates and returns new repository
func NewCityRepository(db *gorm.DB) CityRepositoryInterface {
	return &CityRepository{
		db: db,
	}
}

// GetAllCity and return it
func (c *CityRepository) GetAllCities(name *string, limit *int, offset int) *[]City {
	var cities *[]City
	db := c.db

	if name != nil {
		db = db.Where("name Like ?", *name)
	}

	if limit != nil {
		db = db.Limit(*limit)
	}

	db.Offset(offset).Find(&cities)
	return cities
}

// CountryExists and return is Exists
func (b CountryRepository) CountryExists(countryId uint) bool {
	var country Country
	result := b.db.Where("id = ?", countryId).First(&country)
	return result.Error == nil && !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// CityExists and return is Exists
func (b CityRepository) CityExists(cityId uint) bool {
	var city City
	result := b.db.Where("id = ?", cityId).First(&city)
	return result.Error == nil && !errors.Is(result.Error, gorm.ErrRecordNotFound)
}
