package swipe

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestSwipeUseCase_Like(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createSwipeUseCase(db, randUserId)

	randLikerId := uint(rand.Int())
	randLikedId := uint(rand.Int())
	like := mockLike(randLikerId, randLikedId)

	mockedLikeRequest := mockLikeSwipeRequest(like.LikerId, like.LikedId)
	createLike, err := useCase.Like(ctx, "", mockedLikeRequest)
	assert.NoError(t, err, "like creation failed in Swipe use-case")
	assertLike(t, []Like{*like}, []Like{*createLike})

}

func createSwipeUseCase(db *gorm.DB, userId uint) SwipeUseCaseInterface {

	return NewSwipeUseCase(NewSwipeService(NewSwipeRepository(db)), func(ctx context.Context, token string) (uint, error) {
		return userId, nil
	})
}

func mockLikeSwipeRequest(likerId, likedId uint) *LikeRequest {
	return &LikeRequest{
		LikerId: likerId,
		LikedId: likedId,
	}
}
func mockDisableLikeSwipeRequest(likerId, likedId uint) *LikeRequest {
	return &LikeRequest{
		LikerId: likerId,
		LikedId: likedId,
	}
}

func mockDisLikeSwipeRequest(disLikerId, disLikedId uint) *DisLikeRequest {
	return &DisLikeRequest{
		DisLikerId: disLikerId,
		DisLikedId: disLikedId,
	}
}

func mockDisableDisLikeSwipeRequest(disLikerId, disLikedId uint) *DisLikeRequest {
	return &DisLikeRequest{
		DisLikerId: disLikerId,
		DisLikedId: disLikedId,
	}
}

func mockedGetLikeRequestSwipeRequest(likedId uint) *GetLikeRequest {
	return &GetLikeRequest{
		LikedId: likedId,
	}
}

func mockedGetDisLikeRequestSwipeRequest(disLikedId uint) *GetDisLikeRequest {
	return &GetDisLikeRequest{
		DisLikedId: disLikedId,
	}
}
