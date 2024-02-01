package user

import (
	"gorm.io/gorm"
	"time"
)

// User struct defines user entity which is used in database
type User struct {
	Id          uint           `json:"-" gorm:"primaryKey"`
	UUID        string         `json:"uuid" gorm:"uniqueIndex, not null"`
	Username    *string        `json:"username" gorm:"uniqueIndex, not null"`
	Email       string         `json:"email" gorm:"uniqueIndex, not null"`
	FirstName   string         `json:"first_name,omitempty"`
	LastName    string         `json:"last_name,omitempty"`
	Born        time.Time      `json:"born,omitempty"`
	Country     uint           `json:"country"`
	City        uint           `json:"city"`
	Sex         uint           `json:"sex"`
	Password    string         `json:"-"`
	LastLoginAt *time.Time     `json:"last_login_at,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

// UserRegisterRequest defines a struct for user creation operation
type UserRegisterRequest struct {
	Email           string  `json:"email" validate:"required,email"`
	Username        *string `json:"username,omitempty" validate:"required,omitempty,min:3,max=255"`
	Password        string  `json:"password,omitempty" validate:"required,omitempty, min=8, max=255"`
	PasswordConfirm string  `json:"password_confirm,omitempty" validate:"required,omitempty, min=8, max=255, eqfield=Password"`
}

// UserLoginRequest defines a fields for login operation and appropriate validation rules for that too
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required, min=8, max=255"`
}

// UpdateUserRequest defines a struct which contains necessary fields and data for updating user operation
type UpdateUserRequest struct {
	Username *string `json:"username,omitempty" validate:"required,omitempty,min:3,max=255"`
}

// UpdateUserPasswordRequest defines a struct which contains necessary fields and data for updating user operation
type UpdateUserPasswordRequest struct {
	Password        string `json:"password,omitempty" validate:"required,omitempty,min:8,max:255"`
	PasswordConfirm string `json:"password_confirm,omitempty" validate:"required,omitempty,min=8,max=255,eqfield=Password"`
}

// DeleteUserRequest defines a struct which contains necessary fields and data for updating user operation
type DeleteUserRequest struct {
	Username        *string `json:"username,omitempty" validate:"required,omitempty,min:3,max=255"`
	Password        *string `json:"password,omitempty" validate:"required,omitempty,min:8,max:255"`
	PasswordConfirm *string `json:"password_confirm,omitempty" validate:"required,omitempty,min=8,max=255,eqfield=Password"`
}

// AuthResponse defines struct for authentication operation that contains necessary fields for auth operation
type AuthResponse struct {
	Tokens map[string]string `json:"authorization"`
	User   *User             `json:"user"`
}
