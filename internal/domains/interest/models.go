package interest

import (
	"gorm.io/gorm"
	"time"
)

type Interest struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name,omitempty" gorm:"uniqueIndex, not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// InterestCreateRequest defines a struct for user creation operation
type InterestCreateRequest struct {
	Name string `json:"name,omitempty" validate:"omitempty, min=2, max= 2048"`
}

// InterestGetRequest defines a struct for user creation operation
type InterestGetRequest struct {
	Id     *uint   `json:"id,omitempty" validate:"omitempty, numeric, min=1"`
	Name   *string `json:"name,omitempty" validate:"omitempty,min=1, max=256"`
	Limit  *int    `json:"limit" validate:"required,gte=1"`
	Offset int     `json:"offset" validate:"omitempty,min=0"`
}

// InterestUpdateRequest defines a struct for user creation operation
type InterestUpdateRequest struct {
	Id   *uint   `param:"id,omitempty" validate:"omitempty, numeric, min=1"`
	Name *string `json:"name,omitempty" validate:"omitempty,min=1, max=256"`
}

type InterestDeleteRequest struct {
	Id *uint `param:"id,omitempty" validate:"omitempty, numeric, min=1"`
}
