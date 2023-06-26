package swipe

import (
	"Datearn/pkg/userHandler"
	"context"
)

type SwipeUseCase struct {
	sv        SwipeServiceInterface
	decoderFn func(ctx context.Context, token string) (uint, error)
}

func NewSwipeUseCase(sv SwipeServiceInterface, decoderFn func(ctx context.Context, token string) (uint, error)) SwipeUseCaseInterface {
	if decoderFn == nil {
		decoderFn = userHandler.ExtractUserIdFromToken
	}

	return &SwipeUseCase{
		sv:        sv,
		decoderFn: decoderFn,
	}
}

func (s SwipeUseCase) Like(ctx context.Context, token string, request *LikeRequest) (*Like, error) {
	//TODO implement me
	panic("implement me")
}

func (s SwipeUseCase) DisableLike(ctx context.Context, token string, request *LikeRequest) (*Like, error) {
	//TODO implement me
	panic("implement me")
}

func (s SwipeUseCase) DisLike(ctx context.Context, token string, request *DisLikeRequest) (*DisLike, error) {
	//TODO implement me
	panic("implement me")
}

func (s SwipeUseCase) DisableDisLike(ctx context.Context, token string, request *DisLikeRequest) (*DisLike, error) {
	//TODO implement me
	panic("implement me")
}

func (s SwipeUseCase) GetAllLikes(ctx context.Context, token string, request *GetLikeRequest) (*[]Like, error) {
	//TODO implement me
	panic("implement me")
}
