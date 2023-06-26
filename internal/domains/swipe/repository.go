package swipe

import "gorm.io/gorm"

type SwipeRepository struct {
	db *gorm.DB
}

func NewSwipeRepository(db *gorm.DB) SwipeRepositoryInterface {
	return &SwipeRepository{
		db: db,
	}
}

func (s SwipeRepository) Like(likerId, likedId uint) (*Like, error) {
	like := &Like{
		LikerId: likerId,
		LikedId: likedId,
	}
	result := s.db.Create(like)
	return like, result.Error
}

func (s SwipeRepository) DisableLike(likerId, likedId uint) (*Like, error) {
	//TODO implement me
	panic("implement me")
}

func (s SwipeRepository) DisLike(likerId, likedId uint) (*DisLike, error) {
	//TODO implement me
	panic("implement me")
}

func (s SwipeRepository) DisableDisLike(likerId, likedId uint) (*DisLike, error) {
	//TODO implement me
	panic("implement me")
}

func (s SwipeRepository) GetAllLikes(likedId uint, limit *int, offset int) *[]Like {
	//TODO implement me
	panic("implement me")
}
