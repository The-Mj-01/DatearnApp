package bio

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
	"time"
)

// TestBioRepository_GetBioByUserId functionality
func TestBioRepository_GetBioByUserId(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createBioRepo(db)
	bios := mockAndInsertBio(db, 0, 0, 0, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBioByUserId(bios[0].UserId)
	assertBioEquality(t, []Bio{bios[0]}, []Bio{*fetchedBio})

	randId := rand.Int()
	_, err = repo.GetBioById(uint(randId))
	assert.Error(t, err, "Fetching wrong bio from db failed ! it should throw an error")
}

// TestBioRepository_GetBioById functionality
func TestBioRepository_GetBioById(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createBioRepo(db)
	bios := mockAndInsertBio(db, 0, 0, 0, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBioById(bios[0].Id)
	assertBioEquality(t, bios, []Bio{*fetchedBio})

	randId := rand.Int()
	_, err = repo.GetBioById(uint(randId))
	assert.Error(t, err, "Fetching wrong bio from db failed ! it should throw an error")
}

// TestBioRepository_GetBioByCountry functionality
func TestBioRepository_GetBioByCountry(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createBioRepo(db)

	randId := uint(rand.Int())

	bios := mockAndInsertBio(db, randId, 0, 0, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBatchesBioByCountry(randId)
	assertBioEquality(t, bios, *fetchedBio)

}

// TestBioRepository_GetBioByCity functionality
func TestBioRepository_GetBioByCity(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createBioRepo(db)

	randId := uint(rand.Int())

	bios := mockAndInsertBio(db, 0, randId, 0, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBatchesBioByCity(randId)
	assertBioEquality(t, bios, *fetchedBio)

}

// TestBioRepository_GetBioBySex functionality
func TestBioRepository_GetBioBySex(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createBioRepo(db)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	bios := mockAndInsertBio(db, 0, 0, sexs[0].Id, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBatchesBioBySex(sexs[0].Id)
	assertBioEquality(t, bios, *fetchedBio)

}

func TestBioRepository_SexExists(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createBioRepo(db)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	exists := repo.SexExists(sexs[0].Id)
	assert.True(t, exists, "checking sex existence from repository failed")

	exists = repo.SexExists(uint(rand.Int()))
	assert.False(t, exists, "checking sex existence from repository failed")
}

// TestBioRepository_GetBioByBornAfter functionality
func TestBioRepository_GetBioByBornAfter(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createBioRepo(db)

	bios := mockAndInsertBio(db, 0, 0, 0, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBatchesBioByBornAfter(bios[0].Born)
	assertBioEquality(t, bios, *fetchedBio)

}

// TestBioRepository_GetBioByCountryCitySex functionality
func TestBioRepository_GetBioByCountryCitySex(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createBioRepo(db)

	randCountryId := uint(rand.Int())
	randCityId := uint(rand.Int())

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	bios := mockAndInsertBio(db, randCountryId, randCityId, sexs[0].Id, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBatchesBioByCountryCitySex(randCountryId, randCityId, sexs[0].Id)
	assertBioEquality(t, bios, *fetchedBio)

}

// TestBioRepository_GetBatchesBioByCountryCitySexBornAfterDate functionality
func TestBioRepository_GetBatchesBioByCountryCitySexBornAfterDate(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createBioRepo(db)

	randCountryId := uint(rand.Int())
	randCityId := uint(rand.Int())

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	bios := mockAndInsertBio(db, randCountryId, randCityId, sexs[0].Id, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBatchesBioByCountryCitySexBornAfterDate(randCountryId, randCityId, sexs[0].Id, bios[0].Born)
	assertBioEquality(t, bios, *fetchedBio)
}

// TestBioRepository_CreateBio functionality
func TestBioRepository_CreateBio(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createBioRepo(db)

	randCountryId := uint(rand.Int())
	randCityId := uint(rand.Int())

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	socials := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, socials)

	mockedBio := mockBio(randCountryId, randCityId, sexs[0].Id, 1)

	createdBio, err := repo.CreateBio(mockedBio)
	defer destructCreatedObjects(db, []Bio{*createdBio})

	assert.NoError(t, err, "Bio creation in repository failed")

}

// TestBioRepository_UpdateBio functionality
func TestBioRepository_UpdateBio(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createBioRepo(db)

	randCountryId := uint(rand.Int())
	randCityId := uint(rand.Int())

	sexs := mockAndInsertSex(db, 2)
	defer destructCreatedObjects(db, sexs)

	socials := mockAndInsertSex(db, 2)
	defer destructCreatedObjects(db, socials)

	oldBio := mockAndInsertBio(db, randCountryId, randCityId, sexs[0].Id, 1)
	defer destructCreatedObjects(db, oldBio)
	newBorn := oldBio[0].Born + 10000

	randNewCountryId := uint(rand.Int())
	randNewCityId := uint(rand.Int())

	newBio := Bio{

		Description: "salam",
		Country:     randNewCountryId,
		City:        randNewCityId,
		Sex:         sexs[1].Id,
		Born:        newBorn,
	}

	_, err = repo.UpdateBio(&oldBio[0], &newBio)
	assert.NoError(t, err, "Bio Update operation failed")

	fetchBio := new(Bio)
	db.Where("id = ?", oldBio[0].Id).First(fetchBio)

	assert.Equal(t, oldBio[0].Id, fetchBio.Id, "Bio Update operation failed")
	assert.Equal(t, newBio.Description, fetchBio.Description, "Bio Update operation failed")
	assert.Equal(t, newBio.Country, fetchBio.Country, "Bio Update operation failed")
	assert.Equal(t, newBio.City, fetchBio.City, "Bio Update operation failed")
	assert.Equal(t, newBio.Sex, fetchBio.Sex, "Bio Update operation failed")
	assert.Equal(t, newBio.Born, fetchBio.Born, "Bio Update operation failed")
}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Bio{}, Sex{}, SocialMedia{})
	return db, err
}

// createBioRepo for testing purpose and return it
func createBioRepo(db *gorm.DB) BioRepositoryInterface {
	return NewBioRepository(db)
}

// mockAndInsertSex in database for testing purpose
func mockAndInsertSex(db *gorm.DB, count int) []Sex {
	sexs := make([]Sex, 0, count)
	i := 0
	for {
		tmpSex := mockSex()

		res := db.Create(tmpSex)
		if res.Error != nil {
			continue
		}

		sexs = append(sexs, *tmpSex)
		i += 1

		if i == count {
			break
		}
	}
	return sexs
}

// mockSex object and return it
func mockSex() *Sex {

	return &Sex{

		Title: "male",
	}
}

// mockAndInsertBio in database for testing purpose
func mockAndInsertBio(db *gorm.DB, countryId, cityId, sexId, count uint) []Bio {
	bios := make([]Bio, 0, count)

	for i := 1; i <= int(count); i++ {
		tmpBio := mockBio(countryId, cityId, sexId, uint(i))

		res := db.Create(tmpBio)
		if res.Error != nil {
			continue
		}

		bios = append(bios, *tmpBio)

		if i == int(count) {
			break
		}
	}
	return bios
}

// mockBio object and return it
func mockBio(countryId, cityId, sexId, userId uint) *Bio {
	return &Bio{
		UserId:      userId,
		Description: "this is a description.",
		Country:     countryId,
		City:        cityId,
		Sex:         sexId,
		Born:        time.Now().Unix(),
	}
}

// destructCreatedObjects that are created for testing purpose
func destructCreatedObjects[T Bio | Sex | SocialMedia](db *gorm.DB, records []T) {
	for _, record := range records {
		db.Unscoped().Delete(record)
	}
}

// assertBioEquality to see whether they are equal or not
func assertBioEquality(t *testing.T, mockedBio, fetchedBio []Bio) {
	for index := range mockedBio {
		assert.Equal(t, mockedBio[index].Id, fetchedBio[index].Id)
		assert.Equal(t, mockedBio[index].Country, fetchedBio[index].Country)
		assert.Equal(t, mockedBio[index].City, fetchedBio[index].City)
		assert.Equal(t, mockedBio[index].Sex, fetchedBio[index].Sex)
		assert.Equal(t, mockedBio[index].Born, fetchedBio[index].Born)
	}
}

//// assertCountries check whether they are equal or not
//func assertCountries(t *testing.T, createdCountry, fetchedCountries []Country) {
//	for index := range createdCountry {
//		assert.Equal(t, createdCountry[index].Id, fetchedCountries[index].Id, "Countries Repository test: Ids are not equal")
//		assert.Equal(t, createdCountry[index].Name, fetchedCountries[index].Name, "Countries Repository test: titles are not equal")
//
//	}
//}
//
//// assertCities check whether they are equal or not
//func assertCities(t *testing.T, createdCity, fetchedCities []City) {
//	for index := range createdCity {
//		assert.Equal(t, createdCity[index].Id, fetchedCities[index].Id, "Cities Repository test: Ids are not equal")
//		assert.Equal(t, createdCity[index].Name, fetchedCities[index].Name, "Cities Repository test: titles are not equal")
//
//	}
//}
