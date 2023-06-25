package socialMedia

import "context"

type SocialMediaRepositoryInterface interface {
	GetAllSocialMedia(id *uint, name *string, limit *int, offset int) *[]SocialMedia
	CreateSocialMedia(name string) (*SocialMedia, error)
	UpdateSocialMedia(oldSocial, newSocial *SocialMedia) (*SocialMedia, error)
}

type SocialMediaServiceInterface interface {
	GetAllSocialMedia(id *uint, name *string, limit *int, offset int) (*[]SocialMedia, error)
	CreateSocialMedia(name string) (*SocialMedia, error)
	UpdateSocialMedia(id uint, name string) (*SocialMedia, error)
}

type SocialMediaUseCaseInterface interface {
	GetAllSocialMedia(ctx context.Context, token string, request *SocialMediaGetRequest) (*[]SocialMedia, error)
	CreateSocialMedia(ctx context.Context, token string, request *SocialMediaCreateRequest) (*SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, token string, request *SocialMediaUpdateRequest) (*SocialMedia, error)
}
