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

func createSwipeService(db *gorm.DB) SwipeServiceInterface {
	return NewSwipeService(NewSwipeRepository(db))
}
