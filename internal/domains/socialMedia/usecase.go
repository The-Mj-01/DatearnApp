package socialMedia

import (
	"Datearn/pkg/advancedError"
	"Datearn/pkg/userHandler"
	"context"
)

type SocialMediaUseCase struct {
	sv        SocialMediaServiceInterface
	decoderFn func(ctx context.Context, token string) (uint, error)
}

func NewSocialMediaUseCase(sv SocialMediaServiceInterface, decoderFn func(ctx context.Context, token string) (uint, error)) SocialMediaUseCaseInterface {
	if decoderFn == nil {
		decoderFn = userHandler.ExtractUserIdFromToken
	}

	return &SocialMediaUseCase{
		sv:        sv,
		decoderFn: decoderFn,
	}
}

func (s *SocialMediaUseCase) GetAllSocialMedia(ctx context.Context, token string, request *SocialMediaGetRequest) (*[]SocialMedia, error) {
	_, err := s.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return s.sv.GetAllSocialMedia(request.Id, request.Name, request.Limit, request.Offset)
}

func (s *SocialMediaUseCase) CreateSocialMedia(ctx context.Context, token string, request *SocialMediaCreateRequest) (*SocialMedia, error) {
	_, err := s.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return s.sv.CreateSocialMedia(request.Name)
}
