package bio

import "gorm.io/gorm"

func createService(db *gorm.DB) BioServiceInterface {
	return NewService(NewRepository(db))
}
