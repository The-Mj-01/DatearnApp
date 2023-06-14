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

func TestUserUseCase_UpdateUserName(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.WithValue(context.Background(), ContextValueIpKey, "192.168.1.1")
	usecase := createUseCase(db)
	sv := createService(db)
	mockedUser := mockUser()

	createdUser, err := sv.CreateUser(mockedUser.Username, mockedUser.Email, mockedUser.Password)
	defer destructCreatedObjects(db, []User{*createdUser})

	mockedJwt, err := mockUserJwtTk(ctx, createdUser.UUID, "192.168.1.1")
	assert.NoError(t, err, "Mocking Jwt failed")

	mockedUpdateUserRequest := mockUpdateUserRequest()

	result, err := usecase.UpdateUserName(ctx, mockedJwt, mockedUpdateUserRequest)
	assert.NoError(t, err, "Updating user name failed")
	assert.NotEqual(t, *mockedUser.Username, *result.Username, "Updating user name failed")
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

// mockUpdateUserRequest for testing functionality
func mockUpdateUserRequest() *UpdateUserRequest {
	newUsername := "NewtesingNaaaame"
	newPass := "newtestingPass"
	return &UpdateUserRequest{
		Username:        &newUsername,
		Password:        &newPass,
		PasswordConfirm: &newPass,
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
