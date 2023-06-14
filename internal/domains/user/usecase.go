package user

import (
	"Datearn/pkg/hasher"
	"Datearn/pkg/tokenizer"
	"context"
)

// ContextValueIpKey determines requested key for extracting ip from context
const ContextValueIpKey string = "REQUESTER_IP"

// UserUseCase is a struct which satisfies user use case interface functionalities
type UserUseCase struct {
	sv UserServiceInterface
}

// NewUserUseCase and return it
func NewUserUseCase(sv UserServiceInterface) UserUseCaseInterface {
	return &UserUseCase{
		sv: sv,
	}
}

// Register user in system and generate token for it and then return it
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

// UpdateUserPass for a user that is already exists
func (u *UserUseCase) UpdateUserPass(ctx context.Context, token string, request *UpdateUserRequest) (*User, error) {
	//TODO implement me
	panic("implement me")
}

// UpdateUserName for a user that is already exists
func (u *UserUseCase) UpdateUserName(ctx context.Context, token string, request *UpdateUserRequest) (*User, error) {
	userId, err := u.getUserID(ctx, token)
	if err != nil {
		return nil, AuthSomethingWrong
	}

	user, err := u.sv.UpdateUser(userId, request.Username, nil)
	if err != nil {
		return nil, err
	}

	return user, nil
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

// getUserId decodes token and returns userId form it's UUID
func (u *UserUseCase) getUserID(ctx context.Context, token string) (uint, error) {
	uuid, err := u.getUserUUID(ctx, token)

	fetchedUser, err := u.sv.GetUserByUUID(uuid)
	if err != nil {
		return 0, err
	}

	return fetchedUser.Id, nil
}

// getUserUUID decodes token and returns userId form it's UUID
func (u *UserUseCase) getUserUUID(ctx context.Context, token string) (string, error) {
	tokenGenerator := tokenizer.CreateTokenizer(ctx)
	tkInfo, err := tokenGenerator.TokenInfo(token)
	if err != nil {
		return "", AuthSomethingWrong
	}

	return tkInfo.UUID, nil
}
