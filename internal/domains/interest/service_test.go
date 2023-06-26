package interest

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestInterestService_GetAllInterest(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createInterestService(db)

	_, err = sv.GetAllInterest(nil, nil, nil, 0)
	assert.Error(t, err, "Expected interest not found error")
	assert.ErrorIs(t, err, InterestNotFound, "Expected interest not found error")
}

func TestInterestService_CreateInterest(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createInterestService(db)

	interest := mockInterest()

	createdInterest, err := sv.CreateInterest(interest.Name)
	defer destructCreatedObjects(db, []Interest{*createdInterest})

	assert.NoError(t, err, "interest service bio creation failed")
	assert.Equal(t, interest.Name, createdInterest.Name)

}

func TestInterestService_UpdateInterest(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createInterestService(db)

	oldInterest := mockAndInsertInterest(db, 1)
	defer destructCreatedObjects(db, oldInterest)

	newInterest := &Interest{
		Id:   oldInterest[0].Id,
		Name: "Twitter",
	}

	wrongInterest := &Interest{
		Name: "",
	}

	_, err = sv.UpdateInterest(&newInterest.Id, wrongInterest.Name)
	assert.Error(t, err, "Interest service update functionality failed")
	assert.ErrorIs(t, err, NameNotFound, "Interest service update functionality failed")

	updatedInterest, err := sv.UpdateInterest(&newInterest.Id, newInterest.Name)

	assert.NoError(t, err, "Interest service update user failed")
	assert.Equal(t, newInterest.Id, updatedInterest.Id, "Interest service update bio failed")
	assert.Equal(t, newInterest.Name, updatedInterest.Name, "Interest service update bio failed")

}

func TestInterestService_DeleteInterest(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createInterestService(db)
	interest := mockAndInsertInterest(db, 1)
	defer destructCreatedObjects(db, interest)

	deletedUser, err := service.DeleteInterest(&interest[0].Id)

	assertInterest(t, []Interest{*deletedUser}, []Interest{interest[0]})

	_, err = service.DeleteInterest(&interest[0].Id)
	assert.Error(t, err, "Interest service user creation failed")
	assert.ErrorIs(t, err, InterestNotFound, "Interest service user creation failed")

}

func createInterestService(db *gorm.DB) InterestServiceInterface {
	return NewInterestService(NewInterestRepository(db))
}
