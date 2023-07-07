package image

import "gorm.io/gorm"

func createImageService(db *gorm.DB) ImageServiceInterface {
	return NewImageService(NewImageRepository(db))
}
