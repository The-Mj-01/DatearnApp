package interest

import (
	"context"
	"gorm.io/gorm"
)

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
