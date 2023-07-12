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
		db = db.Where("imageable_type = ?", *imageableType)
	}

	if imageableId != nil {
		db = db.Where("imageable_id = ?", *imageableId)
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
	img := &Image{
		Name:          name,
		ImageableId:   imageableId,
		ImageableType: imageableType,
		Path:          path,
	}
	result := i.db.Create(img)
	return img, result.Error
}

func (i *ImageRepository) UpdateImage(oldImage, newImage *Image) (*Image, error) {
	if newImage.Name != "" {
		oldImage.Name = newImage.Name
	}
	if newImage.Path != "" {
		oldImage.Path = newImage.Path
	}
	if newImage.ImageableId != 0 {
		oldImage.ImageableId = newImage.ImageableId
	}
	if newImage.ImageableType != "" {
		oldImage.ImageableType = newImage.ImageableType
	}

	result := i.db.Save(oldImage)

	return oldImage, result.Error
}

func (i *ImageRepository) DeleteImage(image *Image) (*Image, error) {
	result := i.db.Delete(image)

	return image, result.Error
}
