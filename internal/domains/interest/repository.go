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

func (s *InterestRepository) CreateInterest(name string) (*Interest, error) {
	interest := &Interest{
		Name: name,
	}
	result := s.db.Create(interest)
	return interest, result.Error
}

func (s *InterestRepository) UpdateInterest(oldInterest, newInterest *Interest) (*Interest, error) {
	if newInterest.Name != "" {
		oldInterest.Name = newInterest.Name
	}

	result := s.db.Save(oldInterest)

	return oldInterest, result.Error
}

func (s *InterestRepository) DeleteInterest(interest *Interest) (*Interest, error) {
	result := s.db.Delete(interest)

	return interest, result.Error
}
