package swipe

type SwipeService struct {
	repo SwipeRepositoryInterface
}

func NewSwipeService(repo SwipeRepositoryInterface) SwipeServiceInterface {
	return &SwipeService{
		repo: repo,
	}
}

func (s *SwipeService) Like(likerId, likedId uint) (*Like, error) {
	return s.repo.Like(likerId, likedId)
}

func (s *SwipeService) DisableLike(likerId, likedId uint) (*Like, error) {
	var like *Like

	like = &Like{
		LikerId: likerId,
		LikedId: likedId,
	}
	return s.repo.DisableLike(like)
}

func (s *SwipeService) DisLike(disLikerId, disLikedId uint) (*DisLike, error) {
	return s.repo.DisLike(disLikerId, disLikedId)
}

func (s *SwipeService) DisableDisLike(disLikerId, disLikedId uint) (*DisLike, error) {
	var disLike *DisLike

	disLike = &DisLike{
		DisLikerId: disLikerId,
		DisLikedId: disLikedId,
	}
	return s.repo.DisableDisLike(disLike)
}

func (s *SwipeService) GetAllLikes(likedId uint, limit *int, offset int) (*[]Like, error) {
	//TODO implement me
	panic("implement me")
}
