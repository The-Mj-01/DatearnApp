package swipe

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestSwipeRepository_Like(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createSwipeRepo(db)

	randLikerId := uint(rand.Int())
	randLikedId := uint(rand.Int())

	like := mockLike(randLikerId, randLikedId)

	likedSwipe, err := repo.Like(like.LikerId, like.LikedId)
	assert.NoError(t, err, "like swipe operation failed")

	assertLike(t, []Like{*like}, []Like{*likedSwipe})

}

func TestSwipeRepository_DisableLike(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createSwipeRepo(db)

	randLikerId := []uint{uint(rand.Int())}
	randLikedId := []uint{uint(rand.Int())}

	like := mockAndInsertLike(db, randLikerId, randLikedId, 1)

	disableLikedSwipe, err := repo.DisableLike(&like[0])
	assert.NoError(t, err, "like swipe operation failed")

	assertLike(t, like, []Like{*disableLikedSwipe})

	fetchUser := new(Like)
	result := db.Where("liker_id = ?", like[0].LikerId).Where("liked_id = ?", like[0].LikedId).First(fetchUser)

	assert.Error(t, result.Error, "Like Delete operation failed")

}

func TestSwipeRepository_DisLike(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createSwipeRepo(db)

	randDisLikerId := uint(rand.Int())
	randDisLikedId := uint(rand.Int())

	disLike := mockDisLike(randDisLikerId, randDisLikedId)

	disLikedSwipe, err := repo.DisLike(disLike.DisLikerId, disLike.DisLikedId)
	assert.NoError(t, err, "disLike swipe  operation failed")

	assertDisLike(t, []DisLike{*disLike}, []DisLike{*disLikedSwipe})
}

func TestSwipeRepository_DisableDisLike(t *testing.T) {

	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createSwipeRepo(db)

	randDisLikerId := []uint{uint(rand.Int())}
	randDisLikedId := []uint{uint(rand.Int())}

	disLike := mockAndInsertDisLike(db, randDisLikerId, randDisLikedId, 1)

	disableDisLikedSwipe, err := repo.DisableDisLike(&disLike[0])
	assert.NoError(t, err, "disLike swipe operation failed")

	assertDisLike(t, disLike, []DisLike{*disableDisLikedSwipe})

	fetchUser := new(DisLike)
	result := db.Where("disLiker_id = ?", disLike[0].DisLikerId).Where("disLiked_id = ?", disLike[0].DisLikedId).First(fetchUser)

	assert.Error(t, result.Error, "DisLike Delete operation failed")
}

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
func mockAndInsertLike(db *gorm.DB, likerId, likedId []uint, count int) []Like {
	interest := make([]Like, 0, count)
	i := 0
	for {
		tmpLike := mockLike(likerId[i], likedId[i])

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
func mockLike(likerId, likedId uint) *Like {
	return &Like{
		LikerId: likerId,
		LikedId: likedId,
	}
}

// mockAndInsertDisLike in database for testing purpose
func mockAndInsertDisLike(db *gorm.DB, likerId, likedId []uint, count int) []DisLike {
	interest := make([]DisLike, 0, count)
	i := 0
	for {
		tmpDisLike := mockDisLike(likerId[i], likedId[i])

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
func mockDisLike(likerId, likedId uint) *DisLike {
	return &DisLike{
		DisLikerId: likerId,
		DisLikedId: likedId,
	}
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
