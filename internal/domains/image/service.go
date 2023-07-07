package image

type ImageService struct {
	repo ImageRepositoryInterface
}

func NewImageService(repo ImageRepositoryInterface) ImageServiceInterface {
	return &ImageService{
		repo: repo,
	}
}

func (i *ImageService) GetAllImage(id, imageableId *uint, name, imageableType *string, limit *int, offset int) (*[]Image, error) {
	img := i.repo.GetAllImage(id, imageableId, name, imageableType, limit, offset)

	if len(*img) == 0 {
		return nil, ImageNotFound
	}
	return img, nil
}

func (i *ImageService) CreateImage(imageableId uint, name, path, imageableType string) (*Image, error) {
	if imageableId == 0 {
		return nil, ImageableIdNotFound
	}
	if name == "" {
		return nil, NameNotFound
	}
	if imageableType == "" {
		return nil, ImageableTypeNotFound
	}
	if path == "" {
		return nil, PathNotFound
	}

	return i.repo.CreateImage(imageableId, name, path, imageableType)
}

func (i *ImageService) UpdateImage(id uint, name, path *string) (*Image, error) {

	if name == nil || *name == "" {
		return nil, NameNotFound
	}
	if path == nil || *path == "" {
		return nil, PathNotFound
	}

	newImage := &Image{
		Name: *name,
		Path: *path,
	}

	img, err := i.GetAllImage(&id, nil, nil, nil, nil, 0)

	if err != nil {
		return nil, ImageNotFound
	}

	return i.repo.UpdateImage(&(*img)[0], newImage)
}

func (i *ImageService) DeleteImage(imageId *uint) (*Image, error) {
	//TODO implement me
	panic("implement me")
}
