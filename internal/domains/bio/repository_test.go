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
	bios := mockAndInsertBio(db, 1)
	defer destructCreatedObjects(db, bios)

	fetchedBio, err := repo.GetBioById(bios[0].Id)
	assertUsersEquality(t, fetchedBio, &bios[0])

	randId := rand.Int()
	_, err = repo.GetBioById(uint(randId))
	assert.Error(t, err, "Fetching wrong bio from db failed ! it should throw an error")
}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Bio{}, Country{}, City{}, Sex{})
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

func mockAndInsertBio(db *gorm.DB, count int) []Bio {
	bios := make([]Bio, 0, count)
	i := 0
	for {
		tmpBio := mockBio()

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

func mockBio() *Bio {
	return &Bio{

		SocialMedia: 0,
		Description: "this is a description.",
		Country:     0,
		City:        0,
		Sex:         0,
		Born:        time.Time{},
	}
}

// destructCreatedObjects that are created for testing purpose
func destructCreatedObjects[T Bio | Country | City | Sex](db *gorm.DB, records []T) {
	for _, record := range records {
		db.Unscoped().Delete(record)
	}
}

// assertUsersEquality to see whether they are equal or not
func assertUsersEquality(t *testing.T, fetchedBio, mockedBio *Bio) {
	assert.Equal(t, mockedBio.Id, fetchedBio.Id)
	assert.Equal(t, mockedBio.Country, fetchedBio.Country)
	assert.Equal(t, mockedBio.City, fetchedBio.City)
	assert.Equal(t, mockedBio.Sex, fetchedBio.Sex)
}
