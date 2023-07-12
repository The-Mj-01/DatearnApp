package location

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestCountryService_GetAllCountries(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createCountryService(db)

	_, err = sv.GetAllCountries(nil, nil, 0)
	assert.Error(t, err, "Expected countries not found error")
	assert.ErrorIs(t, err, CountryNotFound, "Expected countries not found error")

}

func TestCityService_GetAllCities(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createCityService(db)

	_, err = sv.GetAllCities(nil, nil, 0)
	assert.Error(t, err, "Expected cities not found error")
	assert.ErrorIs(t, err, CityNotFound, "Expected cities not found error")

}

func createCountryService(db *gorm.DB) CountryServiceInterface {
	return NewCountryService(NewCountryRepository(db))
}

func createCityService(db *gorm.DB) CityServiceInterface {
	return NewCityService(NewCityRepository(db))
}
