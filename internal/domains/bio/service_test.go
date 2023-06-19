package bio

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestBioService_GetBioById functionality
func TestBioService_GetBioById(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createService(db)

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

	service := createService(db)

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

	service := createService(db)

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

	service := createService(db)

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

	service := createService(db)
	bios := mockAndInsertBio(db, 0, 0, 0, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := service.GetBioByBornAfter(bios[0].Born)
	assertBioEquality(t, bios, *fetchedBio)

}

func createService(db *gorm.DB) BioServiceInterface {
	return NewService(NewRepository(db))
}
