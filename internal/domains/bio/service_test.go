package bio

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestBioService_GetBioByUserId functionality
func TestBioService_GetBioByUserId(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createBioService(db)

	countries := mockAndInsertCountry(db, 1)
	defer destructCreatedObjects(db, countries)

	cities := mockAndInsertCity(db, 1)
	defer destructCreatedObjects(db, cities)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	bios := mockAndInsertBio(db, countries[0].Id, cities[0].Id, sexs[0].Id, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := service.GetBioByUserId(bios[0].UserId)
	assertBioEquality(t, bios, []Bio{*fetchedBio})

	randId := rand.Int()
	_, err = service.GetBioByUserId(uint(randId))
	assert.Error(t, err, "Fetching wrong bio from db failed ! it should throw an error")
}

// TestBioService_GetBioById functionality
func TestBioService_GetBioById(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createBioService(db)

	countries := mockAndInsertCountry(db, 1)
	defer destructCreatedObjects(db, countries)

	cities := mockAndInsertCity(db, 1)
	defer destructCreatedObjects(db, cities)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	bios := mockAndInsertBio(db, countries[0].Id, cities[0].Id, sexs[0].Id, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := service.GetBioById(bios[0].Id)
	assertBioEquality(t, bios, []Bio{*fetchedBio})

	randId := rand.Int()
	_, err = service.GetBioById(uint(randId))
	assert.Error(t, err, "Fetching wrong bio from db failed ! it should throw an error")

}

// TestBioService_GetBioByCountry functionality
func TestBioService_GetBioByCountry(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createBioService(db)

	countries := mockAndInsertCountry(db, 1)
	defer destructCreatedObjects(db, countries)

	bios := mockAndInsertBio(db, countries[0].Id, 0, 0, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := service.GetBioByCountry(bios[0].Country)
	assertBioEquality(t, bios, *fetchedBio)
}

// TestBioService_GetBioByCity functionality
func TestBioService_GetBioByCity(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createBioService(db)

	cities := mockAndInsertCity(db, 1)
	defer destructCreatedObjects(db, cities)

	bios := mockAndInsertBio(db, 0, cities[0].Id, 0, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := service.GetBioByCity(bios[0].City)
	assertBioEquality(t, bios, *fetchedBio)
}

// TestBioService_GetBioBySex functionality
func TestBioService_GetBioBySex(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createBioService(db)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	bios := mockAndInsertBio(db, 0, 0, sexs[0].Id, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := service.GetBioBySex(bios[0].Sex)
	assertBioEquality(t, bios, *fetchedBio)
}

// TestBioService_GetBioByBornAfter functionality
func TestBioService_GetBioByBornAfter(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createBioService(db)
	bios := mockAndInsertBio(db, 0, 0, 0, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := service.GetBioByBornAfter(bios[0].Born)
	assertBioEquality(t, bios, *fetchedBio)

}

// TestBioRepository_GetBatchesBioByCountryCitySex functionality
func TestBioRepository_GetBatchesBioByCountryCitySex(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createBioService(db)

	countries := mockAndInsertCountry(db, 1)
	defer destructCreatedObjects(db, countries)

	cities := mockAndInsertCity(db, 1)
	defer destructCreatedObjects(db, cities)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	bios := mockAndInsertBio(db, countries[0].Id, cities[0].Id, sexs[0].Id, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := service.GetBioByCountryCitySex(countries[0].Id, cities[0].Id, sexs[0].Id)
	assertBioEquality(t, bios, *fetchedBio)

}

// TestBioService_GetBioByCountryCitySexBornAfterDate functionality
func TestBioService_GetBioByCountryCitySexBornAfterDate(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createBioService(db)

	countries := mockAndInsertCountry(db, 1)
	defer destructCreatedObjects(db, countries)

	cities := mockAndInsertCity(db, 1)
	defer destructCreatedObjects(db, cities)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	bios := mockAndInsertBio(db, countries[0].Id, cities[0].Id, sexs[0].Id, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := service.GetBioByCountryCitySexBornAfterDate(countries[0].Id, cities[0].Id, sexs[0].Id, bios[0].Born)
	assertBioEquality(t, bios, *fetchedBio)
}

// TestBioService_CreateBio functionality
func TestBioService_CreateBio(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createBioService(db)

	countries := mockAndInsertCountry(db, 1)
	defer destructCreatedObjects(db, countries)

	cities := mockAndInsertCity(db, 1)
	defer destructCreatedObjects(db, cities)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	mockedBio := mockBio(countries[0].Id, cities[0].Id, sexs[0].Id, 1)

	createdBio, err := service.CreateBio(mockedBio.Description, mockedBio.UserId, mockedBio.Country, mockedBio.City, mockedBio.Sex, mockedBio.Born)
	defer destructCreatedObjects(db, []Bio{*createdBio})

	assert.NoError(t, err, "Bio service bio creation failed")
	assert.Equal(t, mockedBio.Country, createdBio.Country)
	assert.Equal(t, mockedBio.City, createdBio.City)
	assert.Equal(t, mockedBio.Sex, createdBio.Sex)
	assert.Equal(t, mockedBio.Born, createdBio.Born)
}

func TestBioService_UpdateBio(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createBioService(db)

	countries := mockAndInsertCountry(db, 2)
	defer destructCreatedObjects(db, countries)

	cities := mockAndInsertCity(db, 2)
	defer destructCreatedObjects(db, cities)

	sexs := mockAndInsertSex(db, 2)
	defer destructCreatedObjects(db, sexs)

	oldBio := mockAndInsertBio(db, countries[0].Id, cities[0].Id, sexs[0].Id, 1)
	defer destructCreatedObjects(db, oldBio)

	newBorn := oldBio[0].Born + 10000
	randId := uint(rand.Int())
	newBio := Bio{
		UserId:      1,
		Description: "salam",
		Country:     countries[1].Id,
		City:        cities[1].Id,
		Sex:         sexs[1].Id,
		Born:        newBorn,
	}
	wrongBioData := Bio{
		UserId:      2,
		Description: "",
		Country:     randId,
		City:        randId,
		Sex:         randId,
		Born:        0,
	}

	_, err = service.UpdateBio(wrongBioData.UserId, newBio.Description, newBio.Country, newBio.City, newBio.Sex, newBio.Born)
	assert.Error(t, err, "Bio service update functionality failed")
	assert.ErrorIs(t, err, YouAreNotAllowed, "Bio service update functionality failed")

	_, err = service.UpdateBio(newBio.UserId, wrongBioData.Description, newBio.Country, newBio.City, newBio.Sex, newBio.Born)
	assert.Error(t, err, "Bio service update functionality failed")
	assert.ErrorIs(t, err, DescripitonNotFound, "Bio service update functionality failed")

	_, err = service.UpdateBio(newBio.UserId, newBio.Description, wrongBioData.Country, newBio.City, newBio.Sex, newBio.Born)
	assert.Error(t, err, "Bio service update functionality failed")
	assert.ErrorIs(t, err, CountryNotFound, "Bio service update functionality failed")

	_, err = service.UpdateBio(newBio.UserId, newBio.Description, newBio.Country, wrongBioData.City, newBio.Sex, newBio.Born)
	assert.Error(t, err, "Bio service update functionality failed")
	assert.ErrorIs(t, err, CityNotFound, "Bio service update functionality failed")

	_, err = service.UpdateBio(newBio.UserId, newBio.Description, newBio.Country, newBio.City, wrongBioData.Sex, newBio.Born)
	assert.Error(t, err, "Bio service update functionality failed")
	assert.ErrorIs(t, err, SexNotFound, "Bio service update functionality failed")

	_, err = service.UpdateBio(newBio.UserId, newBio.Description, newBio.Country, newBio.City, newBio.Sex, wrongBioData.Born)
	assert.Error(t, err, "Bio service update functionality failed")
	assert.ErrorIs(t, err, BornNotFound, "Bio service update functionality failed")

	updatedBio, err := service.UpdateBio(newBio.UserId, newBio.Description, newBio.Country, newBio.City, newBio.Sex, newBio.Born)

	assert.NoError(t, err, "User service update user failed")
	assert.Equal(t, newBio.Description, updatedBio.Description, "Bio service update bio failed")
	assert.Equal(t, newBio.Country, updatedBio.Country, "Bio service update bio failed")
	assert.Equal(t, newBio.City, updatedBio.City, "Bio service update bio failed")
	assert.Equal(t, newBio.Sex, updatedBio.Sex, "Bio service update bio failed")
	assert.Equal(t, newBio.Born, updatedBio.Born, "Bio service update bio failed")

}

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

func createBioService(db *gorm.DB) BioServiceInterface {
	return NewBioService(NewBioRepository(db))
}

func createCountryService(db *gorm.DB) CountryServiceInterface {
	return NewCountryService(NewCountryRepository(db))
}

func createCityService(db *gorm.DB) CityServiceInterface {
	return NewCityService(NewCityRepository(db))
}
