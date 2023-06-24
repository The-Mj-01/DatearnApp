package socialMedia

import "time"

type SocialMedia struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name,omitempty" gorm:"uniqueIndex, not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// SocialMediaCreateRequest defines a struct for user creation operation
type SocialMediaCreateRequest struct {
	Id   uint   `json:"id,omitempty" validate:"omitempty, numeric, min=1"`
	Name string `json:"name,omitempty" validate:"omitempty, min=2, max= 2048"`
}
