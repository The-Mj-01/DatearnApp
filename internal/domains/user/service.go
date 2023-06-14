package user

import (
	"Datearn/pkg/advancedError"
	"Datearn/pkg/hasher"
	"Datearn/pkg/uuid"
)

type UserService struct {
	repo UserRepositoryInterface
}

func NewService(repo UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUserById(id uint) (*User, error) {
	return u.repo.GetUserById(id)

}

func (u *UserService) GetUserByUUID(uuid string) (*User, error) {
	return u.repo.GetUserByUUID(uuid)
}

func (u *UserService) GetUserByEmail(email string) (*User, error) {
	return u.repo.GetUserByEmail(email)
}

func (u *UserService) CreateUser(username *string, email, password string) (*User, error) {
	isExist := u.repo.UserExists(email)
	if isExist {
		return nil, UserAlreadyExists
	}

	hashedPassword, err := hasher.Make(password)
	if err != nil {
		return nil, advancedError.New(err, "user hashing password failed")
	}

	generatedUUID, err := uuid.GenerateUUId()
	if err != nil {
		return nil, advancedError.New(err, "user generate uuid failed")
	}

	user := &User{
		UUID:     generatedUUID,
		Email:    email,
		Username: username,
		Password: hashedPassword,
	}

	return u.repo.CreateUser(user)

}

func (u *UserService) UpdateUser(userId uint, username, password *string) (*User, error) {
	user, err := u.repo.GetUserById(userId)
	if err != nil {
		return nil, UserDoesntExists
	}

	if password != nil {
		password, err = u.makePasswordUseAble(password)
	}

	if err != nil {
		return nil, advancedError.New(err, "Cannot hash password")
	}

	return u.repo.UpdateUser(user, username, password)
}

func (u *UserService) DeleteUser(userId uint, password *string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

// makePasswordUseAble for in update operation
func (u *UserService) makePasswordUseAble(password *string) (*string, error) {
	hashedPass, err := hasher.Make(*password)
	if err != nil {
		return nil, err
	}

	return &hashedPass, nil
}
