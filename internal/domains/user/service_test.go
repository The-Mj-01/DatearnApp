package user

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

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

func createService(db *gorm.DB) UserServiceInterface {
	return NewService(NewRepository(db))
}
