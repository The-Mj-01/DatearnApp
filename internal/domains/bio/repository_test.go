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
	assertBioEquality(t, &bios[0], fetchedBio)

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

	bios := mockAndInsertBio(db, countries[0].Id, 0, 0, 0, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBioByCountry(countries[0].Id)
	assertBioEquality(t, &bios[0], fetchedBio)

	randId := rand.Int()
	_, err = repo.GetBioByCountry(uint(randId))
	assert.Error(t, err, "Fetching wrong bio from db failed ! it should throw an error")
}

func TestBioRepository_GetBioByCity(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)

	cities := mockAndInsertCity(db, 1)
	defer destructCreatedObjects(db, cities)

	bios := mockAndInsertBio(db, 0, cities[0].Id, 0, 0, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBioByCity(cities[0].Id)
	assertBioEquality(t, &bios[0], fetchedBio)

	randId := rand.Int()
	_, err = repo.GetBioByCity(uint(randId))
	assert.Error(t, err, "Fetching wrong bio from db failed ! it should throw an error")

}

func TestBioRepository_GetBioBySex(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	bios := mockAndInsertBio(db, 0, 0, sexs[0].Id, 0, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBioBySex(sexs[0].Id)
	assertBioEquality(t, &bios[0], fetchedBio)

	randId := rand.Int()
	_, err = repo.GetBioBySex(uint(randId))
	assert.Error(t, err, "Fetching wrong bio from db failed ! it should throw an error")
}

func TestBioRepository_GetBioByBornAfter(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)

	bios := mockAndInsertBio(db, 0, 0, 0, 0, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBioByBornAfter(bios[0].Born)
	assertBioEquality(t, &bios[0], fetchedBio)

	randId := rand.Intn(int(time.Now().UnixNano() - 100000000))
	_, err = repo.GetBioByBornAfter(time.Unix(int64(randId), 0))

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

	bios := mockAndInsertBio(db, countries[0].Id, cities[0].Id, sexs[0].Id, 0, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBioByCountryCitySex(countries[0].Id, cities[0].Id, sexs[0].Id)
	assertBioEquality(t, &bios[0], fetchedBio)

	randId := rand.Int()
	_, err = repo.GetBioBySex(uint(randId))
	assert.Error(t, err, "Fetching wrong bio from db failed ! it should throw an error")
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
	i := 0
	for {
		tmpBio := mockBio(countryId, cityId, sexId, socialMediaId)

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

func mockBio(countryId, cityId, sexId, socialMediaId uint) *Bio {
	return &Bio{

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
func assertBioEquality(t *testing.T, mockedBio, fetchedBio *Bio) {
	assert.Equal(t, mockedBio.Id, fetchedBio.Id)
	assert.Equal(t, mockedBio.Country, fetchedBio.Country)
	assert.Equal(t, mockedBio.City, fetchedBio.City)
	assert.Equal(t, mockedBio.Sex, fetchedBio.Sex)
	assert.Equal(t, mockedBio.SocialMedia, fetchedBio.SocialMedia)
	assert.Equal(t, mockedBio.Born, fetchedBio.Born)
}
