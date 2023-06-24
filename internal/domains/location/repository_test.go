package location

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestCountryRepository_GetAllCountries functionality
func TestCountryRepository_GetAllCountries(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createCountryRepo(db)

	countries := mockAndInsertCountry(db, 5)
	defer destructCreatedObjects(db, countries)

	fetchedCountries := repo.GetAllCountries(nil, nil, 10)
	assert.Equal(t, len(*fetchedCountries), 0, "Fetched countries are not equal")

	limit := 1
	fetchedCountries = repo.GetAllCountries(nil, &limit, 0)
	assert.Equal(t, len(*fetchedCountries), limit, "one Country must be fetched")

	falseTitle := "Test irrelevant country title which not exists"
	fetchedCountries = repo.GetAllCountries(&falseTitle, nil, 0)
	assert.Equal(t, len(*fetchedCountries), 0, "zero Country must be fetched")

	fetchedCountries = repo.GetAllCountries(nil, nil, 0)
	assert.NotZero(t, len(*fetchedCountries), "Zero countries fetched")
	assert.Equal(t, len(*fetchedCountries), 5, "Fetched cities are not equal")
	assertCountries(t, countries, *fetchedCountries)

	fetchedCountries = repo.GetAllCountries(&countries[0].Name, nil, 0)
	assert.NotZero(t, len(*fetchedCountries), "Zero countries fetched")
	assertCountries(t, countries, *fetchedCountries)

}

// TestCityRepository_GetAllCities functionality
func TestCityRepository_GetAllCities(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createCityRepo(db)

	cities := mockAndInsertCity(db, 5)
	defer destructCreatedObjects(db, cities)

	fetchedCities := repo.GetAllCities(nil, nil, 10)
	assert.Equal(t, len(*fetchedCities), 0, "Fetched cities are not equal")

	limit := 1
	fetchedCities = repo.GetAllCities(nil, &limit, 0)
	assert.Equal(t, len(*fetchedCities), limit, "one City must be fetched")

	falseTitle := "Test irrelevant city title which not exists"
	fetchedCities = repo.GetAllCities(&falseTitle, nil, 0)
	assert.Equal(t, len(*fetchedCities), 0, "zero City must be fetched")

	fetchedCities = repo.GetAllCities(nil, nil, 0)
	assert.NotZero(t, len(*fetchedCities), "Zero cities fetched")
	assert.Equal(t, len(*fetchedCities), 5, "Fetched cities are not equal")
	assertCities(t, cities, *fetchedCities)
}

// TestBioRepository_CountryExists functionality
func TestBioRepository_CountryExists(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createCountryRepo(db)

	randId := uint(rand.Int())

	exists := repo.CountryExists(randId)
	assert.True(t, exists, "checking country existence from repository failed")

	exists = repo.CountryExists(uint(rand.Int()))
	assert.False(t, exists, "checking country existence from repository failed")

}

// TestBioRepository_CityExists functionality
func TestBioRepository_CityExists(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createCityRepo(db)

	randId := uint(rand.Int())

	exists := repo.CityExists(randId)
	assert.True(t, exists, "checking city existence from repository failed")

	exists = repo.CityExists(uint(rand.Int()))
	assert.False(t, exists, "checking city existence from repository failed")

}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Country{}, City{})
	return db, err
}

// createCountryRepo for testing purpose and return it
func createCountryRepo(db *gorm.DB) CountryRepositoryInterface {
	return NewCountryRepository(db)
}

// createCityRepo for testing purpose and return it
func createCityRepo(db *gorm.DB) CityRepositoryInterface {
	return NewCityRepository(db)
}

// mockAndInsertCountry in database for testing purpose
func mockAndInsertCountry(db *gorm.DB, count int) []Country {
	countries := make([]Country, 0, count)
	i := 0
	for {
		tmpCountry := mockCountry()

		res := db.Create(tmpCountry)
		if res.Error != nil {
			continue
		}

		countries = append(countries, *tmpCountry)
		i += 1

		if i == count {
			break
		}
	}
	return countries
}

// mockCountry object and return it
func mockCountry() *Country {
	return &Country{
		Name: "Iran",
	}
}

// mockAndInsertCity in database for testing purpose
func mockAndInsertCity(db *gorm.DB, count int) []City {
	cities := make([]City, 0, count)
	i := 0
	for {
		tmpCity := mockCity()

		res := db.Create(tmpCity)
		if res.Error != nil {
			continue
		}

		cities = append(cities, *tmpCity)
		i += 1

		if i == count {
			break
		}
	}
	return cities
}

// mockCity object and return it
func mockCity() *City {
	return &City{

		Name: "Tehran",
	}
}

// destructCreatedObjects that are created for testing purpose
func destructCreatedObjects[T Country | City](db *gorm.DB, records []T) {
	for _, record := range records {
		db.Unscoped().Delete(record)
	}
}

// assertCountries check whether they are equal or not
func assertCountries(t *testing.T, createdCountry, fetchedCountries []Country) {
	for index := range createdCountry {
		assert.Equal(t, createdCountry[index].Id, fetchedCountries[index].Id, "Countries Repository test: Ids are not equal")
		assert.Equal(t, createdCountry[index].Name, fetchedCountries[index].Name, "Countries Repository test: titles are not equal")

	}
}

// assertCities check whether they are equal or not
func assertCities(t *testing.T, createdCity, fetchedCities []City) {
	for index := range createdCity {
		assert.Equal(t, createdCity[index].Id, fetchedCities[index].Id, "Cities Repository test: Ids are not equal")
		assert.Equal(t, createdCity[index].Name, fetchedCities[index].Name, "Cities Repository test: titles are not equal")

	}
}
