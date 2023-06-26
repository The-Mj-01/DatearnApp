package swipe

type SwipeService struct {
	repo SwipeRepositoryInterface
}

func NewSwipeService(repo SwipeRepositoryInterface) SwipeServiceInterface {
	return &SwipeService{
		repo: repo,
	}
}

func (s SwipeService) Like(likerId, likedId uint) (*Like, error) {
	//TODO implement me
	panic("implement me")
}

func (s SwipeService) DisableLike(likerId, likedId uint) (*Like, error) {
	//TODO implement me
	panic("implement me")
}

func (s SwipeService) DisLike(likerId, likedId uint) (*DisLike, error) {
	//TODO implement me
	panic("implement me")
}

func (s SwipeService) DisableDisLike(likerId, likedId uint) (*DisLike, error) {
	//TODO implement me
	panic("implement me")
}

func (s SwipeService) GetAllLikes(likedId uint, limit *int, offset int) (*[]Like, error) {
	//TODO implement me
	panic("implement me")
}
