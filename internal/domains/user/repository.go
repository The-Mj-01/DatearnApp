package user

import (
	"errors"
	"gorm.io/gorm"
)

// UserRepository which implements repository interface
type UserRepository struct {
	db *gorm.DB
}

// NewRepository instantiates and returns new repository
func NewRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}

// GetUserById and return it
func (u *UserRepository) GetUserById(id uint) (*User, error) {
	var user User
	result := u.db.Where("id = ?", id).First(&user)
	return &user, result.Error
}

// GetUserByUUID and return it
func (u *UserRepository) GetUserByUUID(uuid string) (*User, error) {
	var user User
	result := u.db.Where("UUID = ?", uuid).First(&user)
	return &user, result.Error
}

// GetUserByEmail and return it
func (u *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	result := u.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

// UserExists checks whether use with this phone already exists or not
func (u *UserRepository) UserExists(email string) bool {
	user := new(User)

	result := u.db.Where("email = ?", email).First(user)

	return result.Error == nil && !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// CreateUser that doesn't already exist in database
func (u *UserRepository) CreateUser(user *User) (*User, error) {
	result := u.db.Create(user)
	return user, result.Error
}

// UpdateUser that already exists in database
func (u *UserRepository) UpdateUser(user *User, username, password *string) (*User, error) {
	if username != nil {
		user.Username = username
	}

	if password != nil {
		user.Password = *password
	}

	result := u.db.Save(user)

	return user, result.Error
}

// DeleteUser that does already exist in database
func (u *UserRepository) DeleteUser(user *User) (*User, error) {

	result := u.db.Delete(user)

	return user, result.Error
}
