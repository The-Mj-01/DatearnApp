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

func (i InterestRepository) GetAllInterest(id *uint, name *string, limit *int, offset int) *[]Interest {
	//TODO implement me
	panic("implement me")
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
