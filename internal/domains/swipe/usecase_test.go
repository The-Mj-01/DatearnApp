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

func TestSwipeUseCase_DisableLike(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createSwipeUseCase(db, randUserId)

	randLikerId := []uint{uint(rand.Int())}
	randLikedId := []uint{uint(rand.Int())}
	mockedLike := mockAndInsertLike(db, randLikerId, randLikedId, 1)

	mockedDisableLikeRequest := mockDisableLikeSwipeRequest(mockedLike[0].LikerId, mockedLike[0].LikedId)

	disableLike, err := useCase.DisableLike(ctx, "", mockedDisableLikeRequest)
	assert.NoError(t, err, "like creation failed in Swipe use-case")

	assertLike(t, mockedLike, []Like{*disableLike})

}

func TestSwipeUseCase_DisLike(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createSwipeUseCase(db, randUserId)

	randDisLikerId := uint(rand.Int())
	randDisLikedId := uint(rand.Int())
	disLike := mockDisLike(randDisLikerId, randDisLikedId)

	mockedDisLikeRequest := mockDisLikeSwipeRequest(disLike.DisLikerId, disLike.DisLikedId)
	createDisLike, err := useCase.DisLike(ctx, "", mockedDisLikeRequest)
	assert.NoError(t, err, "disLike creation failed in Swipe use-case")
	assertDisLike(t, []DisLike{*disLike}, []DisLike{*createDisLike})

}

func TestSwipeUseCase_DisableDisLike(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createSwipeUseCase(db, randUserId)

	randDisLikerId := []uint{uint(rand.Int())}
	randDisLikedId := []uint{uint(rand.Int())}
	mockedDisLike := mockAndInsertDisLike(db, randDisLikerId, randDisLikedId, 1)

	mockedDisableDisLikeRequest := mockDisableDisLikeSwipeRequest(mockedDisLike[0].DisLikerId, mockedDisLike[0].DisLikedId)

	disableDisLike, err := useCase.DisableDisLike(ctx, "", mockedDisableDisLikeRequest)
	assert.NoError(t, err, "like creation failed in Swipe use-case")

	assertDisLike(t, mockedDisLike, []DisLike{*disableDisLike})
}

func TestSwipeUseCase_GetAllLikes(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createSwipeUseCase(db, randUserId)

	randLikerId := []uint{uint(rand.Int())}
	randLikedId := []uint{uint(rand.Int())}
	mockedLike := mockAndInsertLike(db, randLikerId, randLikedId, 1)

	mockedGetLikeRequest := mockedGetLikeRequestSwipeRequest(mockedLike[0].LikedId)

	disableLike, err := useCase.GetAllLikes(ctx, "", mockedGetLikeRequest)
	assert.NoError(t, err, "like creation failed in Swipe use-case")

	assertLike(t, mockedLike, *disableLike)

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
