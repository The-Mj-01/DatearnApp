package swipe

import (
	"gorm.io/gorm"
)

type SwipeRepository struct {
	db *gorm.DB
}

func NewSwipeRepository(db *gorm.DB) SwipeRepositoryInterface {
	return &SwipeRepository{
		db: db,
	}
}

func (s *SwipeRepository) Like(likerId, likedId uint) (*Like, error) {
	like := &Like{
		LikerId: likerId,
		LikedId: likedId,
	}
	result := s.db.Create(like)
	return like, result.Error
}

func (s *SwipeRepository) DisableLike(like *Like) (*Like, error) {
	result := s.db.Delete(like)
	return like, result.Error
}

func (s *SwipeRepository) DisLike(disLikerId, disLikedId uint) (*DisLike, error) {
	disLike := &DisLike{
		DisLikerId: disLikerId,
		DisLikedId: disLikedId,
	}
	result := s.db.Create(disLike)
	return disLike, result.Error
}

func (s *SwipeRepository) DisableDisLike(disLike *DisLike) (*DisLike, error) {
	result := s.db.Delete(disLike)
	return disLike, result.Error
}

func (s *SwipeRepository) GetAllLikes(likedId *uint, limit *int, offset int) *[]Like {
	var like *[]Like

	db := s.db

	if likedId != nil {
		db = db.Where("liked_id LIKE ?", *likedId)
	}

	if limit != nil {
		db = db.Limit(*limit)
	}

	db.Offset(offset).Find(&like)
	return like
}
