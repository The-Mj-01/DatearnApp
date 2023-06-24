package socialMedia

import "gorm.io/gorm"

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
	//TODO implement me
	panic("implement me")
}
