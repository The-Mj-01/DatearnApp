package interest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestInterestUseCase_GetAllInterest(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createInterestUseCase(db, randUserId)

	mockedInterest := mockAndInsertInterest(db, 1)

	assert.Equal(t, len(mockedInterest), 1, "Mocking products failed")

	mockedGetInterestRequest := mockGetInterestRequest(&mockedInterest[0].Id, &mockedInterest[0].Name, nil, 0)

	fetchedInterest, err := useCase.GetAllInterest(ctx, "", mockedGetInterestRequest)
	assert.NotNil(t, fetchedInterest)
	assertInterest(t, mockedInterest, *fetchedInterest)

}

func TestInterestUseCase_CreateInterest(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createInterestUseCase(db, randUserId)

	interest := mockInterest()

	mockedRequest := mockWriteInterestRequest(interest.Name)
	result, err := useCase.CreateInterest(ctx, "", mockedRequest)
	defer destructCreatedObjects(db, []Interest{*result})
	assert.NoError(t, err, "Interest creation failed in address use-case")
	assert.Equal(t, result.Name, mockedRequest.Name, "Interest creation failed in bio use-case")

}

func TestInterestUseCase_UpdateInterest(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createInterestUseCase(db, randUserId)

	oldInterest := mockAndInsertInterest(db, 1)
	defer destructCreatedObjects(db, oldInterest)

	newName := "Twitter"
	mockedEditRequest := mockEditInterestRequest(&oldInterest[0].Id, &newName)
	editedInterest, err := useCase.UpdateInterest(ctx, "", mockedEditRequest)
	defer destructCreatedObjects(db, []Interest{*editedInterest})
	assert.NoError(t, err, "Interest use-case update functionality failed")

	assert.Equal(t, *mockedEditRequest.Name, editedInterest.Name, "Interest use-case update functionality failed")
}

func TestInterestUseCase_DeleteInterest(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createInterestUseCase(db, randUserId)

	mockedInterest := mockAndInsertInterest(db, 1)
	defer destructCreatedObjects(db, mockedInterest)

	mockedDeleteRequest := mockDeleteInterestRequest(&mockedInterest[0].Id)

	deletedInterest, err := useCase.DeleteInterest(ctx, "", mockedDeleteRequest)
	assert.NoError(t, err, "Deleting user name failed")

	assertInterest(t, mockedInterest, []Interest{*deletedInterest})

}

func createInterestUseCase(db *gorm.DB, userId uint) InterestUseCaseInterface {
	return NewInterestUseCase(NewInterestService(NewInterestRepository(db)), func(ctx context.Context, token string) (uint, error) {
		return userId, nil
	})
}

func mockGetInterestRequest(id *uint, name *string, limit *int, offset int) *InterestGetRequest {
	return &InterestGetRequest{
		Id:     id,
		Name:   name,
		Limit:  limit,
		Offset: offset,
	}
}

func mockWriteInterestRequest(name string) *InterestCreateRequest {
	return &InterestCreateRequest{
		Name: name,
	}
}

func mockEditInterestRequest(id *uint, name *string) *InterestUpdateRequest {
	return &InterestUpdateRequest{
		Id:   id,
		Name: name,
	}
}

func mockDeleteInterestRequest(id *uint) *InterestDeleteRequest { //Todo: *uint --> uint
	return &InterestDeleteRequest{
		Id: id,
	}
}
