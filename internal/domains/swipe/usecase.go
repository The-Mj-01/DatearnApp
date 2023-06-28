package swipe

import (
	"Datearn/pkg/advancedError"
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
	_, err := s.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return s.sv.Like(request.LikerId, request.LikedId)
}

func (s SwipeUseCase) DisableLike(ctx context.Context, token string, request *LikeRequest) (*Like, error) {
	_, err := s.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return s.sv.DisableLike(request.LikerId, request.LikedId)
}

func (s SwipeUseCase) DisLike(ctx context.Context, token string, request *DisLikeRequest) (*DisLike, error) {
	_, err := s.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return s.sv.DisLike(request.DisLikerId, request.DisLikedId)
}

func (s SwipeUseCase) DisableDisLike(ctx context.Context, token string, request *DisLikeRequest) (*DisLike, error) {
	_, err := s.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return s.sv.DisableDisLike(request.DisLikerId, request.DisLikedId)
}

func (s SwipeUseCase) GetAllLikes(ctx context.Context, token string, request *GetLikeRequest) (*[]Like, error) {
	//TODO implement me
	panic("implement me")
}
