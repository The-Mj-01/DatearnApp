package bio

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
	"time"
)

func TestBioRepository_GetBioById(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)
	bios := mockAndInsertBio(db, 0, 0, 0, 0, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBioById(bios[0].Id)
	assertBioEquality(t, bios, []Bio{*fetchedBio})

	randId := rand.Int()
	_, err = repo.GetBioById(uint(randId))
	assert.Error(t, err, "Fetching wrong bio from db failed ! it should throw an error")
}

func TestBioRepository_GetBioByCountry(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)

	countries := mockAndInsertCountry(db, 1)
	defer destructCreatedObjects(db, countries)

	bios := mockAndInsertBio(db, countries[0].Id, 0, 0, 0, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBatchesBioByCountry(countries[0].Id)
	assertBioEquality(t, bios, *fetchedBio)

}

func TestBioRepository_GetBioByCity(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)

	cities := mockAndInsertCity(db, 1)
	defer destructCreatedObjects(db, cities)

	bios := mockAndInsertBio(db, 0, cities[0].Id, 0, 0, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBatchesBioByCity(cities[0].Id)
	assertBioEquality(t, bios, *fetchedBio)

}

func TestBioRepository_GetBioBySex(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	bios := mockAndInsertBio(db, 0, 0, sexs[0].Id, 0, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBatchesBioBySex(sexs[0].Id)
	assertBioEquality(t, bios, *fetchedBio)

}

func TestBioRepository_GetBioByBornAfter(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)

	bios := mockAndInsertBio(db, 0, 0, 0, 0, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBatchesBioByBornAfter(bios[0].Born)
	assertBioEquality(t, bios, *fetchedBio)

}

func TestBioRepository_GetBioByCountryCitySex(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)

	countries := mockAndInsertCountry(db, 1)
	defer destructCreatedObjects(db, countries)

	cities := mockAndInsertCity(db, 1)
	defer destructCreatedObjects(db, cities)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	bios := mockAndInsertBio(db, countries[0].Id, cities[0].Id, sexs[0].Id, 0, 10)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBatchesBioByCountryCitySex(countries[0].Id, cities[0].Id, sexs[0].Id)
	assertBioEquality(t, bios, *fetchedBio)

}

func TestBioRepository_GetBatchesBioByCountryCitySexBornAfterDate(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)

	countries := mockAndInsertCountry(db, 1)
	defer destructCreatedObjects(db, countries)

	cities := mockAndInsertCity(db, 1)
	defer destructCreatedObjects(db, cities)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	bios := mockAndInsertBio(db, countries[0].Id, cities[0].Id, sexs[0].Id, 0, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBatchesBioByCountryCitySexBornAfterDate(countries[0].Id, cities[0].Id, sexs[0].Id, bios[0].Born)
	assertBioEquality(t, bios, *fetchedBio)
}

func TestBioRepository_CreateBio(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)

	countries := mockAndInsertCountry(db, 1)
	defer destructCreatedObjects(db, countries)

	cities := mockAndInsertCity(db, 1)
	defer destructCreatedObjects(db, cities)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	socials := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, socials)

	mockedBio := mockBio(countries[0].Id, countries[0].Id, cities[0].Id, sexs[0].Id, socials[0].Id)

	createdBio, err := repo.CreateBio(mockedBio)
	defer destructCreatedObjects(db, []Bio{*createdBio})

	assert.NoError(t, err, "Bio creation in repository failed")

}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Bio{}, Country{}, City{}, Sex{}, SocialMedia{})
	return db, err
}

// createRepo for testing purpose and return it
func createRepo(db *gorm.DB) BioRepositoryInterface {
	return NewRepository(db)
}

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

func mockCountry() *Country {
	return &Country{
		Name: "Iran",
	}
}

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

func mockCity() *City {
	return &City{

		Name: "Tehran",
	}
}

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

func mockSex() *Sex {

	return &Sex{

		Title: "male",
	}
}

func mockAndInsertSocialMedia(db *gorm.DB, count int) []SocialMedia {
	medias := make([]SocialMedia, 0, count)
	i := 0
	for {
		tmpSocialMedia := mockSocialMedia()

		res := db.Create(tmpSocialMedia)
		if res.Error != nil {
			continue
		}

		medias = append(medias, *tmpSocialMedia)
		i += 1

		if i == count {
			break
		}
	}
	return medias
}

func mockSocialMedia() *SocialMedia {

	return &SocialMedia{

		Name:          "instagram",
		SocialMediaId: "@ali",
	}
}

func mockAndInsertBio(db *gorm.DB, countryId, cityId, sexId, socialMediaId uint, count int) []Bio {
	bios := make([]Bio, 0, count)

	for i := 0; i < count; i++ {
		tmpBio := mockBio(countryId, cityId, sexId, socialMediaId, uint(i))

		res := db.Create(tmpBio)
		if res.Error != nil {
			continue
		}

		bios = append(bios, *tmpBio)
		i += 1

		if i == count {
			break
		}
	}
	return bios
}

func mockBio(countryId, cityId, sexId, socialMediaId, index uint) *Bio {
	return &Bio{
		Id:          index,
		SocialMedia: socialMediaId,
		Description: "this is a description.",
		Country:     countryId,
		City:        cityId,
		Sex:         sexId,
		Born:        time.Time{},
	}
}

// destructCreatedObjects that are created for testing purpose
func destructCreatedObjects[T Bio | Country | City | Sex | SocialMedia](db *gorm.DB, records []T) {
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
		assert.Equal(t, mockedBio[index].SocialMedia, fetchedBio[index].SocialMedia)
		assert.Equal(t, mockedBio[index].Born, fetchedBio[index].Born)
	}
}
