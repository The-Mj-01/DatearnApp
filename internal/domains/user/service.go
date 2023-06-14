package user

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

func (u *UserService) UserExists(email string) bool {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) CreateUser(username *string, email, password string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) UpdateUser(userId uint, username, password *string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) DeleteUser(userId uint, password *string) (*User, error) {
	//TODO implement me
	panic("implement me")
}
