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

	disableLike, err := sv.DisableLike(mockedLike[0].LikerId, mockedLike[0].LikedId)
	assert.NoError(t, err, "Disabale like failed")
	assertLike(t, mockedLike, []Like{*disableLike})
}

func createSwipeService(db *gorm.DB) SwipeServiceInterface {
	return NewSwipeService(NewSwipeRepository(db))
}
