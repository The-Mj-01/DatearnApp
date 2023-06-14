package user

import (
	"Datearn/pkg/tokenizer"
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestUserUseCase_Register(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	useCase := createUseCase(db)
	registerReq := mockRegisterRequest()
	ctx := context.WithValue(context.Background(), ContextValueIpKey, "192.168.1.1")

	result, err := useCase.Register(ctx, registerReq)
	defer destructCreatedObjects(db, []User{*result.User})

	assert.NoError(t, err, "register failed in user use-case")
	assert.NotNil(t, result.Tokens, "register failed in user use-case")
	assert.NotNil(t, result.User, "register failed in user use-case")
	assert.Equal(t, *result.User.Username, *registerReq.Username, "register failed in user use-case")
}

func TestUserUseCase_Login(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.WithValue(context.Background(), ContextValueIpKey, "192.168.1.1")
	usecase := createUseCase(db)
	sv := createService(db)

	mockedLoginReq := mockLoginRequest()
	createdUser, err := sv.CreateUser(nil, mockedLoginReq.Email, mockedLoginReq.Password)
	defer destructCreatedObjects(db, []User{*createdUser})

	assert.NoError(t, err, "Use creation in test user use-case failed")

	result, err := usecase.Login(ctx, mockedLoginReq)
	assert.NoError(t, err, "login failed in user use-case")
	assert.NotNil(t, result.Tokens, "login failed in user use-case")
	assert.NotNil(t, result.User, "login failed in user use-case")
	assertUsersEquality(t, result.User, createdUser)
}

// createUseCase and return it for testing purpose
func createUseCase(db *gorm.DB) UserUseCaseInterface {
	return NewUserUseCase(NewService(NewRepository(db)))
}

// mockLoginRequest and return it for login operation
func mockLoginRequest() *UserLoginRequest {
	return &UserLoginRequest{
		Email:    "example@gmail.com",
		Password: "1234567879",
	}
}

// mockRegisterRequest and return it for register operation
func mockRegisterRequest() *UserRegisterRequest {
	name := "KingApr"
	return &UserRegisterRequest{
		Email:           "example@gmail.com",
		Password:        "1234567879",
		PasswordConfirm: "1234567879",
		Username:        &name,
	}
}

// mockUserJwtTk and return it
func mockUserJwtTk(ctx context.Context, uuid, ip string) (string, error) {
	tokenGenerator := tokenizer.CreateTokenizer(ctx)
	jwtTk, err := tokenGenerator.New(uuid, ip)
	if err != nil {
		return "", err
	}

	return jwtTk["token"], nil
}
