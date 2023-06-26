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

func mockDeleteInterestRequest(id *uint) *InterestDeleteRequest {
	return &InterestDeleteRequest{
		Id: id,
	}
}
