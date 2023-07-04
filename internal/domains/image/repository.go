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

func (i *ImageRepository) GetAllImage(id, imageableId *uint, name, imageableType *string, limit *int, offset int) *[]Image {
	var interest *[]Image

	db := i.db

	if name != nil {
		db = db.Where("name LIKE ?", *name)
	}

	if imageableType != nil {
		db = db.Where("imageableType = ?", *imageableType)
	}

	if imageableId != nil {
		db = db.Where("imageableId = ?", *imageableId)
	}

	if id != nil {
		db = db.Where("id = ?", *id)
	}

	if limit != nil {
		db = db.Limit(*limit)
	}

	db.Offset(offset).Find(&interest)
	return interest
}

func (i *ImageRepository) CreateImage(imageableId uint, name, path, imageableType string) (*Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i *ImageRepository) UpdateImage(oldImage, newImage *Image) (*Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i *ImageRepository) DeleteImage(image *Image) (*Image, error) {
	//TODO implement me
	panic("implement me")
}
