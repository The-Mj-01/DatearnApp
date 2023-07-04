package image

import "gorm.io/gorm"

type ImageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) ImageRepositoryInterface {
	return &ImageRepository{
		db: db,
	}
}

func (i ImageRepository) GetAllImage(id, imageableId *uint, name, imageableType *string, limit *int, offset int) *[]Image {
	//TODO implement me
	panic("implement me")
}

func (i ImageRepository) CreateImage(imageableId uint, name, path, imageableType string) (*Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i ImageRepository) UpdateImage(oldImage, newImage *Image) (*Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i ImageRepository) DeleteImage(image *Image) (*Image, error) {
	//TODO implement me
	panic("implement me")
}
