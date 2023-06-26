package swipe

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Like{}, DisLike{})
	return db, err
}

// createSwipeRepo for testing purpose and return it
func createSwipeRepo(db *gorm.DB) SwipeRepositoryInterface {
	return NewSwipeRepository(db)
}

// mockAndInsertLike in database for testing purpose
func mockAndInsertLike(db *gorm.DB, count int) []Like {
	interest := make([]Like, 0, count)
	i := 0
	for {
		tmpLike := mockLike()

		res := db.Create(tmpLike)
		if res.Error != nil {
			continue
		}

		interest = append(interest, *tmpLike)
		i += 1

		if i == count {
			break
		}
	}
	return interest
}

// mockLike object and return it
func mockLike() *Like {
	return &Like{}
}

// mockAndInsertDisLike in database for testing purpose
func mockAndInsertDisLike(db *gorm.DB, count int) []DisLike {
	interest := make([]DisLike, 0, count)
	i := 0
	for {
		tmpDisLike := mockDisLike()

		res := db.Create(tmpDisLike)
		if res.Error != nil {
			continue
		}

		interest = append(interest, *tmpDisLike)
		i += 1

		if i == count {
			break
		}
	}
	return interest
}

// mockDisLike object and return it
func mockDisLike() *DisLike {
	return &DisLike{}
}

// assertLike check whether they are equal or not
func assertLike(t *testing.T, createdLike, fetchedLike []Like) {
	for index := range createdLike {
		assert.Equal(t, createdLike[index].LikerId, fetchedLike[index].LikerId, "Like Repository test: Ids are not equal")
		assert.Equal(t, createdLike[index].LikedId, fetchedLike[index].LikedId, "Like Repository test: titles are not equal")

	}
}

// assertDisLike check whether they are equal or not
func assertDisLike(t *testing.T, createdDisLike, fetchedDisLike []DisLike) {
	for index := range createdDisLike {
		assert.Equal(t, createdDisLike[index].DisLikerId, fetchedDisLike[index].DisLikerId, "DisLike Repository test: Ids are not equal")
		assert.Equal(t, createdDisLike[index].DisLikedId, fetchedDisLike[index].DisLikedId, "DisLike Repository test: titles are not equal")

	}
}

// destructCreatedObjects that are created for testing purpose
func destructCreatedObjects[T Like | DisLike](db *gorm.DB, records []T) {
	for _, record := range records {
		db.Unscoped().Delete(record)
	}
}
