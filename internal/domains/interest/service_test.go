package interest

import "gorm.io/gorm"

func createInterestService(db *gorm.DB) InterestServiceInterface {
	return NewInterestService(NewInterestRepository(db))
}
