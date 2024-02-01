package swipe

import (
	"gorm.io/gorm"
	"time"
)

type Like struct {
	LikerId   uint           `json:"liker_id,omitempty" gorm:"primaryKey"`
	LikedId   uint           `json:"liked_id,omitempty" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type DisLike struct {
	DisLikerId uint           `json:"dis_liker_id,omitempty" gorm:"primaryKey"`
	DisLikedId uint           `json:"dis_liked_id,omitempty" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

// LikeRequest defines a struct for user creation operation
type LikeRequest struct {
	LikerId uint `json:"liker_id,omitempty" validate:"omitempty, numeric, min=1"`
	LikedId uint `json:"liked_id,omitempty" validate:"omitempty, numeric, min=1"`
}

// DisLikeRequest defines a struct for user creation operation
type DisLikeRequest struct {
	DisLikerId uint `json:"liker_id,omitempty" validate:"omitempty, numeric, min=1"`
	DisLikedId uint `json:"liked_id,omitempty" validate:"omitempty, numeric, min=1"`
}

// GetLikeRequest defines a struct for user creation operation
type GetLikeRequest struct {
	LikedId uint `json:"liked_id,omitempty" validate:"omitempty, numeric, min=1"`
	Limit   *int `json:"limit" validate:"required,gte=1"`
	Offset  int  `json:"offset" validate:"omitempty,min=0"`
}

// GetLikeRequest defines a struct for user creation operation
type GetDisLikeRequest struct {
	DisLikedId uint `json:"liked_id,omitempty" validate:"omitempty, numeric, min=1"`
}
