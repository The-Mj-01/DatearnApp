package user

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestUserService_GetUserById functionality
func TestUserService_GetUserById(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createService(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	fetchUser, err := service.GetUserById(users[0].Id)
	assertUsersEquality(t, fetchUser, &users[0])

	randId := rand.Int()
	_, err = service.GetUserById(uint(randId))
	assert.Error(t, err, "Fetching wrong user from db failed ! it should throw an error")

}

// TestUserService_GetUserByUUID functionality
func TestUserService_GetUserByUUID(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createService(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	fetchUser, err := service.GetUserByUUID(users[0].UUID)
	assertUsersEquality(t, fetchUser, &users[0])

	randUUID := "test2UUID"
	_, err = service.GetUserByUUID(randUUID)
	assert.Error(t, err, "Fetching wrong user from db failed ! it should throw an error")
}

// TestUserService_GetUserByEmail functionality
func TestUserService_GetUserByEmail(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createService(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	fetchUser, err := service.GetUserByEmail(users[0].Email)
	assertUsersEquality(t, fetchUser, &users[0])

	randEmail := "example2@gmail.com"
	_, err = service.GetUserByEmail(randEmail)
	assert.Error(t, err, "Fetching wrong user from db failed ! it should throw an error")

}

// TestUserService_CreateUser functionality
func TestUserService_CreateUser(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createService(db)
	mockedUser := mockUser()

	createdUser, err := service.CreateUser(mockedUser.Username, mockedUser.Email, mockedUser.Password)
	defer destructCreatedObjects(db, []User{*createdUser})

	assert.NoError(t, err, "User service user creation failed")
	assert.NotEqual(t, createdUser.Password, mockedUser.Password, "User service user creation failed")
	assert.NotZero(t, createdUser.Id, "User service user creation failed")

	_, err = service.CreateUser(mockedUser.Username, mockedUser.Email, mockedUser.Password)
	assert.Error(t, err, "User service user creation failed")
	assert.ErrorIs(t, err, UserAlreadyExists, "User service user creation failed")
}

// TestUserService_UpdateUser functionality
func TestUserService_UpdateUser(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createService(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	newName := "newlyyyname"
	pass := "Password"

	updatedUser, err := service.UpdateUser(users[0].Id, &newName, &pass)
	assert.NoError(t, err, "User service update user failed")
	assert.Equal(t, *updatedUser.Username, newName, "User service update user failed")
	assert.NotEqual(t, pass, updatedUser.Password, "User service update user failed")

	randId := rand.Int()
	_, err = service.UpdateUser(uint(randId), &newName, &pass)
	assert.Error(t, err, "User service update user failed")
	assert.ErrorIs(t, err, UserDoesntExists, "User service update user failed")
}

// TestUserService_DeleteUser functionality
func TestUserService_DeleteUser(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createService(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	deletedUser, err := service.DeleteUser(users[0].Id)

	assertUsersEquality(t, deletedUser, &users[0])

	_, err = service.DeleteUser(users[0].Id)
	assert.Error(t, err, "User service user creation failed")
	assert.ErrorIs(t, err, UserDoesntExists, "User service user creation failed")

}

func createService(db *gorm.DB) UserServiceInterface {
	return NewService(NewRepository(db))
}
