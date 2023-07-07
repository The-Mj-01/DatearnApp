package image

type ImageService struct {
	repo ImageRepositoryInterface
}

func NewImageService(repo ImageRepositoryInterface) ImageServiceInterface {
	return &ImageService{
		repo: repo,
	}
}

func (i ImageService) GetAllImage(id, imageableId *uint, name, imageableType *string, limit *int, offset int) (*[]Image, error) {
	img := i.repo.GetAllImage(id, imageableId, name, imageableType, limit, offset)

	if len(*img) == 0 {
		return nil, ImageNotFound
	}
	return img, nil
}

func (i ImageService) CreateImage(imageableId uint, name, path, imageableType string) (*Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i ImageService) UpdateImage(id *uint, name string) (*Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i ImageService) DeleteImage(imageId *uint) (*Image, error) {
	//TODO implement me
	panic("implement me")
}
