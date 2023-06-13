package user

import "gorm.io/gorm"

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

func (u *UserRepository) GetUserById(id int) (*User, error) {
	var user User
	result := u.db.Where("id = ?", id).First(&user)
	return &user, result.Error
}

func (u *UserRepository) GetUserByUUID(uuid string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) GetUserByEmail(email string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) UserExists(email string) bool {
	//TODO implement me
	panic("implement me")
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
