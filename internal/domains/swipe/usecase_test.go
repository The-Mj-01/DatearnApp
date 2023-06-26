package swipe

import (
	"context"
	"gorm.io/gorm"
)

func createSwipeUseCase(db *gorm.DB, userId uint, withValidFunc bool) SwipeUseCaseInterface {

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
