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

func (u *UserRepository) GetUserById(id uint) (*User, error) {
	var user User
	result := u.db.Where("id = ?", id).First(&user)
	return &user, result.Error
}

func (u *UserRepository) GetUserByUUID(uuid string) (*User, error) {
	var user User
	result := u.db.Where("UUID = ?", uuid).First(&user)
	return &user, result.Error
}

func (u *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	result := u.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (u *UserRepository) UserExists(email string) bool {
	user := new(User)

	result := u.db.Where("email = ?", email).First(user)

	return result.Error == nil && !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (u *UserRepository) CreateUser(user *User) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) UpdateUser(user *User, username, password *string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) DeleteUser(user *User) (*User, error) {
	//TODO implement me
	panic("implement me")
}
