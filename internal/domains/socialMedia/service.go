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

func (s *SocialMediaService) UpdateSocialMedia(id *uint, name string) (*SocialMedia, error) {
	if name == "" {
		return nil, NameNotFound
	}
	newSocialMedia := &SocialMedia{
		Name: name,
	}

	social, err := s.GetAllSocialMedia(id, nil, nil, 0)

	if err != nil {
		return nil, SocialMediaNotFound
	}

	return s.repo.UpdateSocialMedia(&(*social)[0], newSocialMedia)

}

func (s *SocialMediaService) DeleteSocialMedia(socialId *uint) (*SocialMedia, error) {
	socialMedia := s.repo.GetAllSocialMedia(socialId, nil, nil, 0)
	if len(*socialMedia) == 0 {
		return nil, SocialMediaNotFound
	}

	return s.repo.DeleteSocialMedia(&(*socialMedia)[0])
}
