package socialMedia

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestSocialMediaRepository_GetAllSocialMedia(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createSocialMediaRepo(db)

	social := mockAndInsertSocialMedia(db, 5)
	defer destructCreatedObjects(db, social)

	fetchedSocialMedia := repo.GetAllSocialMedia(nil, nil, nil, 10)
	assert.Equal(t, len(*fetchedSocialMedia), 0, "Fetched social media are not equal")

	limit := 1
	fetchedSocialMedia = repo.GetAllSocialMedia(nil, nil, &limit, 0)
	assert.Equal(t, len(*fetchedSocialMedia), limit, "one Social media must be fetched")

	falseTitle := "Test irrelevant social media title which not exists"
	fetchedSocialMedia = repo.GetAllSocialMedia(nil, &falseTitle, nil, 0)
	assert.Equal(t, len(*fetchedSocialMedia), 0, "zero Social media must be fetched")

	fetchedSocialMedia = repo.GetAllSocialMedia(nil, nil, nil, 0)
	assert.NotZero(t, len(*fetchedSocialMedia), "Zero social media fetched")
	assert.Equal(t, len(*fetchedSocialMedia), 5, "Fetched social media are not equal")
	assertSocialMedia(t, social, *fetchedSocialMedia)

	fetchedSocialMedia = repo.GetAllSocialMedia(nil, &social[0].Name, nil, 0)
	assert.NotZero(t, len(*fetchedSocialMedia), "Zero social media fetched")
	assertSocialMedia(t, social, *fetchedSocialMedia)

}

func TestSocialMediaRepository_CreateSocialMedia(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createSocialMediaRepo(db)

	social := mockSocialMedia()

	createdSocialMedia, err := repo.CreateSocialMedia(social.Name)
	defer destructCreatedObjects(db, []SocialMedia{*createdSocialMedia})

	assert.NoError(t, err, "Bio creation in repository failed")
	assert.Equal(t, social.Name, createdSocialMedia.Name, "SocialMedia Repository test: titles are not equal")

}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(SocialMedia{})
	return db, err
}

// createSocialMediaRepo for testing purpose and return it
func createSocialMediaRepo(db *gorm.DB) SocialMediaRepositoryInterface {
	return NewSocialMediaRepository(db)
}

// mockAndInsertSocialMedia in database for testing purpose
func mockAndInsertSocialMedia(db *gorm.DB, count int) []SocialMedia {
	social := make([]SocialMedia, 0, count)
	i := 0
	for {
		tmpSocialMedia := mockSocialMedia()

		res := db.Create(tmpSocialMedia)
		if res.Error != nil {
			continue
		}

		social = append(social, *tmpSocialMedia)
		i += 1

		if i == count {
			break
		}
	}
	return social
}

// mockSocialMedia object and return it
func mockSocialMedia() *SocialMedia {
	return &SocialMedia{
		Name: "instagram",
	}
}

// assertSocialMedia check whether they are equal or not
func assertSocialMedia(t *testing.T, createdSocialMedia, fetchedSocialMedia []SocialMedia) {
	for index := range createdSocialMedia {
		assert.Equal(t, createdSocialMedia[index].Id, fetchedSocialMedia[index].Id, "SocialMedia Repository test: Ids are not equal")
		assert.Equal(t, createdSocialMedia[index].Name, fetchedSocialMedia[index].Name, "SocialMedia Repository test: titles are not equal")

	}
}

// destructCreatedObjects that are created for testing purpose
func destructCreatedObjects[T SocialMedia](db *gorm.DB, records []T) {
	for _, record := range records {
		db.Unscoped().Delete(record)
	}
}
