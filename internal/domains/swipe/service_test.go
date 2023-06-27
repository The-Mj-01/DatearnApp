package swipe

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestSwipeService_Like(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createSwipeService(db)
	randLikerId := uint(rand.Int())
	randLikedId := uint(rand.Int())
	like := mockLike(randLikerId, randLikedId)

	createdLike, err := sv.Like(like.LikerId, like.LikedId)
	assert.NoError(t, err, "swipe service like creation failed")
	assertLike(t, []Like{*like}, []Like{*createdLike})
}

func TestSwipeService_DisableLike(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createSwipeService(db)

	randLikerId := uint(rand.Int())
	randLikedId := uint(rand.Int())
	mockedLike := mockAndInsertLike(db, []uint{randLikerId}, []uint{randLikedId}, 1)
	defer destructCreatedObjects(db, mockedLike)

	disableLike, err := sv.DisableLike(mockedLike[0].LikerId, mockedLike[0].LikedId)
	assert.NoError(t, err, "Disabale like failed")
	assertLike(t, mockedLike, []Like{*disableLike})
}

func TestSwipeService_DisLike(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createSwipeService(db)
	randDisLikerId := uint(rand.Int())
	randDisLikedId := uint(rand.Int())
	disLike := mockDisLike(randDisLikerId, randDisLikedId)

	createdDisLike, err := sv.DisLike(disLike.DisLikerId, disLike.DisLikedId)
	assert.NoError(t, err, "swipe service disLike creation failed")
	assertDisLike(t, []DisLike{*disLike}, []DisLike{*createdDisLike})
}

func TestSwipeService_DisableDisLike(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createSwipeService(db)

	randDisLikerId := uint(rand.Int())
	randDisLikedId := uint(rand.Int())
	mockedDisLike := mockAndInsertDisLike(db, []uint{randDisLikerId}, []uint{randDisLikedId}, 1)
	defer destructCreatedObjects(db, mockedDisLike)

	disableDisLike, err := sv.DisableDisLike(mockedDisLike[0].DisLikerId, mockedDisLike[0].DisLikedId)
	assert.NoError(t, err, "Disabale like failed")
	assertDisLike(t, mockedDisLike, []DisLike{*disableDisLike})
}

func TestSwipeService_GetAllLikes(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createSwipeService(db)

	randn := uint(rand.Int())
	randLikerId := []uint{uint(rand.Int()), uint(rand.Int())}
	randLikedId := []uint{randn, randn}
	mockedLike := mockAndInsertLike(db, randLikerId, randLikedId, 2)
	defer destructCreatedObjects(db, mockedLike)

	fetchedLike, err := sv.GetAllLikes(&randn, nil, 0)

	assert.NoError(t, err, "Get all likes fetch query failed")

	assertLike(t, mockedLike, *fetchedLike)

}

func createSwipeService(db *gorm.DB) SwipeServiceInterface {
	return NewSwipeService(NewSwipeRepository(db))
}
