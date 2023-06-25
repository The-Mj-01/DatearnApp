package socialMedia

type SocialMediaService struct {
	repo SocialMediaRepositoryInterface
}

func NewSocialMediaService(repo SocialMediaRepositoryInterface) SocialMediaServiceInterface {
	return &SocialMediaService{
		repo: repo,
	}
}

func (s *SocialMediaService) GetAllSocialMedia(id *uint, name *string, limit *int, offset int) (*[]SocialMedia, error) {
	social := s.repo.GetAllSocialMedia(id, name, limit, offset)
	if len(*social) == 0 {
		return nil, SocialMediaNotFound
	}

	return social, nil
}

func (s *SocialMediaService) CreateSocialMedia(name string) (*SocialMedia, error) {
	if name == "" {
		return nil, NameNotFound
	}

	return s.repo.CreateSocialMedia(name)
}

func (s *SocialMediaService) UpdateSocialMedia(id uint, name string) (*SocialMedia, error) {
	//TODO implement me
	panic("implement me")
}
