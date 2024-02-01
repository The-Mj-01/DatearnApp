package socialMedia

import (
	"gorm.io/gorm"
	"time"
)

type SocialMedia struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name,omitempty" gorm:"uniqueIndex, not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// SocialMediaCreateRequest defines a struct for user creation operation
type SocialMediaCreateRequest struct {
	Name string `json:"name,omitempty" validate:"omitempty, min=2, max= 2048"`
}

// SocialMediaGetRequest defines a struct for user creation operation
type SocialMediaGetRequest struct {
	Id     *uint   `json:"id,omitempty" validate:"omitempty, numeric, min=1"`
	Name   *string `json:"name,omitempty" validate:"omitempty,min=1, max=256"`
	Limit  *int    `json:"limit" validate:"required,gte=1"`
	Offset int     `json:"offset" validate:"omitempty,min=0"`
}

// SocialMediaUpdateRequest defines a struct for user creation operation
type SocialMediaUpdateRequest struct {
	Id   *uint   `json:"id,omitempty" validate:"omitempty, numeric, min=1"`
	Name *string `json:"name,omitempty" validate:"omitempty,min=1, max=256"`
}

type SocialMediaDeleteRequest struct {
	Id *uint `json:"id,omitempty" validate:"omitempty, numeric, min=1"`
}
