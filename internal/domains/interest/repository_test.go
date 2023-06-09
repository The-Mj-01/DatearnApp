package interest

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestInterestRepository_GetAllInterest(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createInterestRepo(db)

	interest := mockAndInsertInterest(db, 5)
	defer destructCreatedObjects(db, interest)

	fetchedInterest := repo.GetAllInterest(nil, nil, nil, 10)
	assert.Equal(t, len(*fetchedInterest), 0, "Fetched interest  are not equal")

	limit := 1
	fetchedInterest = repo.GetAllInterest(nil, nil, &limit, 0)
	assert.Equal(t, len(*fetchedInterest), limit, "one interest must be fetched")

	falseTitle := "Test irrelevant interest  title which not exists"
	fetchedInterest = repo.GetAllInterest(nil, &falseTitle, nil, 0)
	assert.Equal(t, len(*fetchedInterest), 0, "zero interest must be fetched")

	fetchedInterest = repo.GetAllInterest(nil, nil, nil, 0)
	assert.NotZero(t, len(*fetchedInterest), "Zero interest  fetched")
	assert.Equal(t, len(*fetchedInterest), 5, "Fetched interest  are not equal")
	assertInterest(t, interest, *fetchedInterest)

	fetchedInterest = repo.GetAllInterest(nil, &interest[0].Name, nil, 0)
	assert.NotZero(t, len(*fetchedInterest), "Zero interest  fetched")
	assertInterest(t, interest, *fetchedInterest)

}

func TestInterestRepository_CreateInterest(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createInterestRepo(db)

	interest := mockInterest()

	createdInterest, err := repo.CreateInterest(interest.Name)
	defer destructCreatedObjects(db, []Interest{*createdInterest})

	assert.NoError(t, err, "Bio creation in repository failed")
	assert.Equal(t, interest.Name, createdInterest.Name, "Interest Repository test: titles are not equal")

}

func TestInterestRepository_UpdateInterest(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createInterestRepo(db)

	oldMockedInterest := mockAndInsertInterest(db, 1)
	defer destructCreatedObjects(db, oldMockedInterest)

	newInterest := &Interest{
		Id:   oldMockedInterest[0].Id,
		Name: "Twitter",
	}

	_, err = repo.UpdateInterest(&oldMockedInterest[0], newInterest)
	assert.NoError(t, err, "Interest Update operation failed")

	fetchInterest := new(Interest)
	db.Where("id = ?", oldMockedInterest[0].Id).First(fetchInterest)

	assert.Equal(t, oldMockedInterest[0].Id, fetchInterest.Id, "Interest Update operation failed")
	assert.Equal(t, newInterest.Name, fetchInterest.Name, "Interest Update operation failed")
}

func TestInterestRepository_DeleteInterest(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed!")

	repo := createInterestRepo(db)
	interest := mockAndInsertInterest(db, 1)
	defer destructCreatedObjects(db, interest)

	deletedInterest, err := repo.DeleteInterest(&interest[0])

	assertInterest(t, []Interest{*deletedInterest}, []Interest{interest[0]})
	fetchUser := new(Interest)
	result := db.Where("id = ?", interest[0].Id).First(fetchUser)

	assert.Error(t, result.Error, "Interest Delete operation failed")
}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Interest{})
	return db, err
}

// createInterestRepo for testing purpose and return it
func createInterestRepo(db *gorm.DB) InterestRepositoryInterface {
	return NewInterestRepository(db)
}

// mockAndInsertInterest in database for testing purpose
func mockAndInsertInterest(db *gorm.DB, count int) []Interest {
	interest := make([]Interest, 0, count)
	i := 0
	for {
		tmpInterest := mockInterest()

		res := db.Create(tmpInterest)
		if res.Error != nil {
			continue
		}

		interest = append(interest, *tmpInterest)
		i += 1

		if i == count {
			break
		}
	}
	return interest
}

// mockInterest object and return it
func mockInterest() *Interest {
	return &Interest{
		Name: "instagram",
	}
}

// assertInterest check whether they are equal or not
func assertInterest(t *testing.T, createdInterest, fetchedInterest []Interest) {
	for index := range createdInterest {
		assert.Equal(t, createdInterest[index].Id, fetchedInterest[index].Id, "Interest Repository test: Ids are not equal")
		assert.Equal(t, createdInterest[index].Name, fetchedInterest[index].Name, "Interest Repository test: titles are not equal")

	}
}

// destructCreatedObjects that are created for testing purpose
func destructCreatedObjects[T Interest](db *gorm.DB, records []T) {
	for _, record := range records {
		db.Unscoped().Delete(record)
	}
}
