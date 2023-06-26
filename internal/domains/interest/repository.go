package interest

import "gorm.io/gorm"

type InterestRepository struct {
	db *gorm.DB
}

func NewInterestRepository(db *gorm.DB) InterestRepositoryInterface {
	return &InterestRepository{
		db: db,
	}
}

func (s *InterestRepository) GetAllInterest(id *uint, name *string, limit *int, offset int) *[]Interest {
	var interest *[]Interest

	db := s.db

	if name != nil {
		db = db.Where("name LIKE ?", *name)
	}

	if limit != nil {
		db = db.Limit(*limit)
	}

	db.Offset(offset).Find(&interest)
	return interest
}

func (i InterestRepository) CreateInterest(name string) (*Interest, error) {
	//TODO implement me
	panic("implement me")
}

func (i InterestRepository) UpdateInterest(oldSocial, newSocial *Interest) (*Interest, error) {
	//TODO implement me
	panic("implement me")
}

func (i InterestRepository) DeleteInterest(socialMedia *Interest) (*Interest, error) {
	//TODO implement me
	panic("implement me")
}