package user

import "context"

// UserRepositoryInterface defines set of methods which every repository who wants to play role as user repo should obey
type UserRepositoryInterface interface {
	GetUserById(id uint) (*User, error)
	GetUserByUUID(uuid string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UserExists(email string) bool
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User, username, password *string) (*User, error)
	DeleteUser(user *User) (*User, error)
}

// UserServiceInterface defines set of methods which every service who wants to play role as user service should obey.
type UserServiceInterface interface {
	GetUserById(id uint) (*User, error)
	GetUserByUUID(uuid string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UserExists(email string) bool
	CreateUser(username *string, email, password string) (*User, error)
	UpdateUser(userId uint, username, password *string) (*User, error)
	DeleteUser(userId uint, password *string) (*User, error)
}

// UserUseCaseInterface defines set of methods which every use case who wants to play role as user use case should obey.
type UserUseCaseInterface interface {
	Register(ctx context.Context, request *UserRegisterRequest) (*AuthResponse, error)
	Login(ctx context.Context, request *UserLoginRequest) (*AuthResponse, error)
	UpdateUserPass(ctx context.Context, request *UpdateUserRequest) (*User, error)
	UpdateUserName(ctx context.Context, request *UpdateUserRequest) (*User, error)
	DeleteUser(ctx context.Context, request *DeleteUserRequest) (*User, error)
}
