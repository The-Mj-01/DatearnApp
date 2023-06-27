package swipe

import "context"

type SwipeRepositoryInterface interface {
	Like(likerId, likedId uint) (*Like, error)
	DisableLike(like *Like) (*Like, error)
	DisLike(likerId, likedId uint) (*DisLike, error)
	DisableDisLike(disLike *DisLike) (*DisLike, error)

	GetAllLikes(likedId *uint, limit *int, offset int) *[]Like
	//GetAllDisLikes(likedId uint, limit *int, offset int) (*[]DisLike)

}

type SwipeServiceInterface interface {
	Like(likerId, likedId uint) (*Like, error)
	DisableLike(likerId, likedId uint) (*Like, error)
	DisLike(likerId, likedId uint) (*DisLike, error)
	DisableDisLike(likerId, likedId uint) (*DisLike, error)

	GetAllLikes(likedId uint, limit *int, offset int) (*[]Like, error)
	//GetAllDisLikes(likedId uint, limit *int, offset int) (*[]DisLike, error)

}

type SwipeUseCaseInterface interface {
	Like(ctx context.Context, token string, request *LikeRequest) (*Like, error)
	DisableLike(ctx context.Context, token string, request *LikeRequest) (*Like, error)
	DisLike(ctx context.Context, token string, request *DisLikeRequest) (*DisLike, error)
	DisableDisLike(ctx context.Context, token string, request *DisLikeRequest) (*DisLike, error)

	GetAllLikes(ctx context.Context, token string, request *GetLikeRequest) (*[]Like, error)
	//GetAllDisLikes(ctx context.Context, token string, request *GetDisLikeRequest) (*[]DisLike, error)

}
