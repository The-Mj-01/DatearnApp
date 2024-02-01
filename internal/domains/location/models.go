package location

import (
	"gorm.io/gorm"
	"time"
)

// Country struct defines user entity which is used in database
type Country struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name,omitempty" gorm:"uniqueIndex, not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// City struct defines user entity which is used in database
type City struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name,omitempty" gorm:"Index, not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type CountryGetrequest struct {
	Name   *string `json:"name,omitempty" validate:"omitempty,min=1, max=256"`
	Limit  *int    `json:"limit" validate:"required,gte=1"`
	Offset int     `json:"offset" validate:"omitempty,min=0"`
}

type CityGetrequest struct {
	Name   *string `json:"name,omitempty" validate:"omitempty,min=1, max= 256"`
	Limit  *int    `json:"limit" validate:"required,gte=1"`
	Offset int     `json:"offset" validate:"omitempty,min=0"`
}
