package socialMedia

import (
	"gorm.io/gorm"
)

type SocialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepositoryInterface {
	return &SocialMediaRepository{
		db: db,
	}
}

func (s *SocialMediaRepository) GetAllSocialMedia(id *uint, name *string, limit *int, offset int) *[]SocialMedia {
	var social *[]SocialMedia

	db := s.db

	if name != nil {
		db = db.Where("name LIKE ?", *name)
	}

	if limit != nil {
		db = db.Limit(*limit)
	}

	db.Offset(offset).Find(&social)
	return social
}

func (s *SocialMediaRepository) CreateSocialMedia(name string) (*SocialMedia, error) {
	social := &SocialMedia{
		Name: name,
	}
	result := s.db.Create(social)
	return social, result.Error
}

func (s *SocialMediaRepository) UpdateSocialMedia(oldSocial, newSocial *SocialMedia) (*SocialMedia, error) {
	if newSocial.Name != "" {
		oldSocial.Name = newSocial.Name
	}

	result := s.db.Save(oldSocial)

	return oldSocial, result.Error
}

func (s *SocialMediaRepository) DeleteSocialMedia(socialMedia *SocialMedia) (*SocialMedia, error) {
	result := s.db.Delete(socialMedia)

	return socialMedia, result.Error
}
