package bio

import "time"

// Bio struct defines user entity which is used in database
type Bio struct {
	Id          uint      `json:"-,omitempty" gorm:"primaryKey"`
	UserId      uint      `json:"userId,omitempty" gorm:"uniqueIndex"`
	SocialMedia uint      `json:"social_media,omitempty"`
	Description string    `json:"description,omitempty" `
	Country     uint      `json:"country"`
	City        uint      `json:"city"`
	Sex         uint      `json:"sex"`
	Born        time.Time `json:"born,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type SocialMedia struct {
	Id            uint      `json:"id" gorm:"primaryKey"`
	Name          uint      `json:"name,omitempty" gorm:"uniqueIndex, not null"`
	SocialMediaId uint      `json:"social_media_id,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

// Country struct defines user entity which is used in database
type Country struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      uint      `json:"name,omitempty" gorm:"uniqueIndex, not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// City struct defines user entity which is used in database
type City struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      uint      `json:"name,omitempty" gorm:"Index, not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// Sex struct defines user entity which is used in database
type Sex struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Title     uint      `json:"title,omitempty" gorm:"uniqueIndex, not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// BioCreateRequest defines a struct for user creation operation
type BioCreateRequest struct {
	UserId      uint      `json:"userId,omitempty" validate:"omitempty, numeric, min=1"`
	SocialMedia uint      `json:"social_media,omitempty" validate:"omitempty,numeric,min=1"`
	Description string    `json:"description,omitempty" validate:"omitempty, min=10, max= 2048"`
	Country     uint      `json:"country" validate:"omitempty,numeric, min=1"`
	City        uint      `json:"city" validate:"omitempty,numeric,min=1"`
	Sex         uint      `json:"sex" validate:"omitempty,numeric,min=1"`
	Born        time.Time `json:"born,omitempty" validate:"omitempty"`
}

type BioUpdateRequest struct {
	SocialMedia uint      `json:"social_media,omitempty" validate:"omitempty,numeric,min=1"`
	Description string    `json:"description,omitempty" validate:"omitempty, min=10, max= 2048"`
	Country     uint      `json:"country" validate:"omitempty,numeric, min=1"`
	City        uint      `json:"city" validate:"omitempty,numeric,min=1"`
	Sex         uint      `json:"sex" validate:"omitempty,numeric,min=1"`
	Born        time.Time `json:"born,omitempty" validate:"omitempty"`
}

type BioGetSingleRequest struct {
	UserId uint `json:"userId,omitempty" validate:"omitempty, numeric, min=1"`
}
