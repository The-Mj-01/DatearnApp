package bio

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
	"time"
)

func TestBioUseCase_WriteBio(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")
	ctx := context.Background()

	randUserId := uint(rand.Int())
	useCase := createUseCase(db, randUserId)
	countries := mockAndInsertCountry(db, 1)
	defer destructCreatedObjects(db, countries)

	cities := mockAndInsertCity(db, 1)
	defer destructCreatedObjects(db, cities)

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	born := time.Now().Unix()
	description := "this is a description!"
	mockedRequest := mockWriteBioRequest(randUserId, countries[0].Id, cities[0].Id, sexs[0].Id, description, born)

	result, err := useCase.WriteBio(ctx, "", mockedRequest)
	assert.NoError(t, err, "Bio creation failed in address use-case")

	assert.Equal(t, result.Country, mockedRequest.Country, "Bio creation failed in bio use-case")
	assert.Equal(t, result.City, mockedRequest.City, "Bio creation failed in bio use-case")
	assert.Equal(t, result.Sex, mockedRequest.Sex, "Bio creation failed in bio use-case")
	assert.Equal(t, result.Born, mockedRequest.Born, "Bio creation failed in bio use-case")
	assert.Equal(t, result.Description, mockedRequest.Description, "Bio creation failed in bio use-case")
	assert.Equal(t, result.UserId, mockedRequest.UserId, "Bio creation failed in bio use-case")
}

func createUseCase(db *gorm.DB, userId uint) BioUseCaseInterface {
	return NewBioUseCase(NewService(NewRepository(db)), func(ctx context.Context, token string) (uint, error) {
		return userId, nil
	})
}

func mockWriteBioRequest(userId, countryId, cityId, sexId uint, description string, born int64) *BioCreateRequest {
	return &BioCreateRequest{
		UserId:      userId,
		Description: description,
		Country:     countryId,
		City:        cityId,
		Sex:         sexId,
		Born:        born,
	}
}
