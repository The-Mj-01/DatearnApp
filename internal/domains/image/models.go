package image

import (
	"image"
	"time"
)

type Image struct {
	Id            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name,omitempty" gorm:"not null"`
	Path          string    `json:"path,omitempty" gorm:"not null"`
	ImageableId   uint      `json:"imageable_id,omitempty" gorm:"not null"`
	ImageableType string    `json:"imageable_type,omitempty" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

// ImageCreateRequest defines a struct for user creation operation
type ImageCreateRequest struct {
	Name          string      `json:"name,omitempty" validate:"omitempty, min=2, max= 2048"`
	ImageableId   uint        `json:"imageable_id,omitempty" validate:"omitempty, min=1"`
	ImageableType string      `json:"imageable_type,omitempty" validate:"omitempty, min=3, max= 256"`
	Img           *image.RGBA `json:"image,omitempty" validate:"omitempty"`
}

// ImageGetRequest defines a struct for user creation operation
type ImageGetRequest struct {
	Id            *uint   `json:"id,omitempty" validate:"omitempty, numeric, min=1"`
	Name          *string `json:"name,omitempty" validate:"omitempty,min=1, max=256"`
	ImageableId   *uint   `json:"imageable_id,omitempty" validate:"omitempty, min=1"`
	ImageableType *string `json:"imageable_type,omitempty" validate:"omitempty, min=3, max= 256"`
	Limit         *int    `json:"limit" validate:"required,gte=1"`
	Offset        int     `json:"offset" validate:"omitempty,min=0"`
}

// ImageUpdateRequest defines a struct for user creation operation
type ImageUpdateRequest struct {
	Id            *uint       `param:"id,omitempty" validate:"omitempty, numeric, min=1"`
	Name          *string     `json:"name,omitempty" validate:"omitempty,min=1, max=256"`
	ImageableId   *uint       `json:"imageable_id,omitempty" validate:"omitempty, min=1"`
	ImageableType *string     `json:"imageable_type,omitempty" validate:"omitempty, min=3, max= 256"`
	Img           *image.RGBA `json:"image,omitempty" validate:"omitempty"`
}

type ImageDeleteRequest struct {
	Id uint `param:"id,omitempty" validate:"omitempty, numeric, min=1"`
}
