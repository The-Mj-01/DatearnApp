package user

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestUserRepository_GetUserById functionality
func TestUserRepository_GetUserById(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	fetchedUser, err := repo.GetUserById(users[0].Id)
	assertUsersEquality(t, fetchedUser, &users[0])

	randId := rand.Int()
	_, err = repo.GetUserById(uint(randId))
	assert.Error(t, err, "Fetching wrong user from db failed ! it should throw an error")
}

// TestUserRepository_GetUserByUUID functionality
func TestUserRepository_GetUserByUUID(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	fetchedUser, err := repo.GetUserByUUID(users[0].UUID)
	assertUsersEquality(t, fetchedUser, &users[0])

	randUUID := "Test2UUid"
	_, err = repo.GetUserByUUID(randUUID)
	assert.Error(t, err, "Fetching wrong user from db failed ! it should throw an error")
}

// TestUserRepository_GetUserByEmail functionality
func TestUserRepository_GetUserByEmail(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	fetchedUser, err := repo.GetUserByEmail(users[0].Email)
	assertUsersEquality(t, fetchedUser, &users[0])

	randEmail := "ali@gmail.com"
	_, err = repo.GetUserByEmail(randEmail)
	assert.Error(t, err, "Fetching wrong user from db failed ! it should throw an error")
}

// TestUserRepository_UserExists functionality
func TestUserRepository_UserExists(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	exists := repo.UserExists(users[0].Email)
	assert.True(t, exists, "checking User existence failed")

	randEmail := "ali@gmail.com"
	exists = repo.UserExists(randEmail)

	assert.False(t, exists, "checking User existence failed")
}

// TestUserRepository_CreateUser functionality
func TestUserRepository_CreateUser(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)
	mockedUser := mockUser()

	createdUser, err := repo.CreateUser(mockedUser)
	defer destructCreatedObjects(db, []User{*createdUser})

	assert.NoError(t, err, "User creation in repository failed")
	assertUsersEquality(t, createdUser, mockedUser)
}

// TestUserRepository_UpdateUser functionality
func TestUserRepository_UpdateUser(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	repo := createRepo(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	newName := "new name"
	newPass := "New pass"

	_, err = repo.UpdateUser(&users[0], &newName, &newPass)
	assert.NoError(t, err, "User Update operation failed")

	fetchUser := new(User)
	db.Where("id = ?", users[0].Id).First(fetchUser)

	assert.Equal(t, *fetchUser.Username, newName, "User Update operation failed")
	assert.Equal(t, fetchUser.Password, newPass, "User Update operation failed")
}

// TestUserRepository_DeleteUser functionality
func TestUserRepository_DeleteUser(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed!")

	repo := createRepo(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	DeletedUser, err := repo.DeleteUser(&users[0])

	assertUsersEquality(t, DeletedUser, &users[0])
}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(User{})
	return db, err
}

// createRepo for testing purpose and return it
func createRepo(db *gorm.DB) UserRepositoryInterface {
	return NewRepository(db)
}

// mockAndInsertUser in database for testing purpose
func mockAndInsertUser(db *gorm.DB, count int) []User {
	users := make([]User, 0, count)
	i := 0
	for {
		tmpUser := mockUser()

		res := db.Create(tmpUser)
		if res.Error != nil {
			continue
		}

		users = append(users, *tmpUser)
		i += 1

		if i == count {
			break
		}
	}
	return users
}

// mockUser object and return it
func mockUser() *User {
	name := "Mohammad"

	return &User{
		UUID:     "testUUID",
		Email:    "example@gmail.com",
		Username: &name,
		Password: "High operative password",
	}
}

// assertUsersEquality to see whether they are equal or not
func assertUsersEquality(t *testing.T, fetchedUser, mockedUser *User) {
	assert.Equal(t, fetchedUser.Id, mockedUser.Id)
	assert.Equal(t, fetchedUser.Username, mockedUser.Username)
	assert.Equal(t, fetchedUser.UUID, mockedUser.UUID)
	assert.Equal(t, fetchedUser.Email, mockedUser.Email)
}

// destructCreatedObjects that are created for testing purpose
func destructCreatedObjects(db *gorm.DB, users []User) {
	for _, user := range users {
		db.Unscoped().Delete(user)
	}
}
