package user

import (
	"Datearn/pkg/hasher"
	"Datearn/pkg/tokenizer"
	"context"
)

const ContextValueIpKey string = "REQUESTER_IP"

type UserUseCase struct {
	sv UserServiceInterface
}

func NewUserUseCase(sv UserServiceInterface) UserUseCaseInterface {
	return &UserUseCase{
		sv: sv,
	}
}

func (u *UserUseCase) Register(ctx context.Context, request *UserRegisterRequest) (*AuthResponse, error) {
	user, err := u.sv.CreateUser(request.Username, request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	requestIp, err := u.extractIP(ctx)
	if err != nil {
		return nil, err
	}

	jwtTk, err := u.generateToken(ctx, requestIp, user.UUID)
	if err != nil {
		return nil, err
	}
	return &AuthResponse{
		Tokens: jwtTk,
		User:   user,
	}, nil
}

// Login user and generate token for them return it
func (u *UserUseCase) Login(ctx context.Context, request *UserLoginRequest) (*AuthResponse, error) {
	user, err := u.sv.GetUserByEmail(request.Email)
	if err != nil {
		return nil, UserDoesntExists
	}

	if !hasher.Check(request.Password, user.Password) {
		return nil, WrongCredentials
	}

	requesterIP, err := u.extractIP(ctx)
	if err != nil {
		return nil, err
	}

	jwtTk, err := u.generateToken(ctx, requesterIP, user.UUID)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Tokens: jwtTk,
		User:   user,
	}, nil
}

func (u *UserUseCase) UpdateUserPass(ctx context.Context, request *UpdateUserRequest) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) UpdateUserName(ctx context.Context, request *UpdateUserRequest) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) DeleteUser(ctx context.Context, request *DeleteUserRequest) (*User, error) {
	//TODO implement me
	panic("implement me")
}

// extractIP from given context
func (u *UserUseCase) extractIP(ctx context.Context) (string, error) {
	requesterIP := ctx.Value(ContextValueIpKey)
	requesterIPStr, ok := requesterIP.(string)
	if !ok || requesterIPStr == "" {
		return "", AuthSomethingWrong
	}

	return requesterIPStr, nil
}

// generateToken with help of tokenizer pkg. it's kind of a wrapper
func (u *UserUseCase) generateToken(ctx context.Context, requestedIp, uuid string) (map[string]string, error) {
	tokenGenerator := tokenizer.CreateTokenizer(ctx)
	jwtTk, err := tokenGenerator.New(uuid, requestedIp)
	if err != nil {
		return nil, AuthSomethingWrong
	}
	return jwtTk, nil
}
