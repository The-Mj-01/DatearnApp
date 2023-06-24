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
	useCase := createBioUseCase(db, randUserId, true)

	randCountryId := uint(rand.Int())
	randCityId := uint(rand.Int())

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	born := time.Now().Unix()
	description := "this is a description!"
	mockedRequest := mockWriteBioRequest(randUserId, randCountryId, randCityId, sexs[0].Id, description, born)

	result, err := useCase.WriteBio(ctx, "", mockedRequest)
	assert.NoError(t, err, "Bio creation failed in address use-case")

	assert.Equal(t, result.Country, mockedRequest.Country, "Bio creation failed in bio use-case")
	assert.Equal(t, result.City, mockedRequest.City, "Bio creation failed in bio use-case")
	assert.Equal(t, result.Sex, mockedRequest.Sex, "Bio creation failed in bio use-case")
	assert.Equal(t, result.Born, mockedRequest.Born, "Bio creation failed in bio use-case")
	assert.Equal(t, result.Description, mockedRequest.Description, "Bio creation failed in bio use-case")
	assert.Equal(t, result.UserId, mockedRequest.UserId, "Bio creation failed in bio use-case")
}

func TestBioUseCase_UpdateBio(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createBioUseCase(db, randUserId, true)

	randCountryId := uint(rand.Int())
	randCityId := uint(rand.Int())

	sexs := mockAndInsertSex(db, 2)
	defer destructCreatedObjects(db, sexs)

	oldBio := mockAndInsertBio(db, randCountryId, randCityId, sexs[0].Id, 1)
	defer destructCreatedObjects(db, oldBio)

	newBorn := oldBio[0].Born + 10000
	newDescription := "salam"
	randNewCountryId := uint(rand.Int())
	randNewCityId := uint(rand.Int())

	mockedEditRequest := mockEditBioRequest(randNewCountryId, randNewCityId, sexs[1].Id, newDescription, newBorn)
	editedBio, err := useCase.UpdateBio(ctx, "", mockedEditRequest)
	assert.NoError(t, err, "Bio use-case update functionality failed")

	assert.Equal(t, mockedEditRequest.Description, editedBio.Description, "UserAddress use-case update functionality failed")
	assert.Equal(t, mockedEditRequest.Country, editedBio.Country, "UserAddress use-case update functionality failed")
	assert.Equal(t, mockedEditRequest.City, editedBio.City, "UserAddress use-case update functionality failed")
	assert.Equal(t, mockedEditRequest.Sex, editedBio.Sex, "UserAddress use-case update functionality failed")
	assert.Equal(t, mockedEditRequest.Born, editedBio.Born, "UserAddress use-case update functionality failed")
}

func TestBioUseCase_GetBio(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createBioUseCase(db, randUserId, true)

	randCountryId := uint(rand.Int())
	randCityId := uint(rand.Int())

	sexs := mockAndInsertSex(db, 1)
	defer destructCreatedObjects(db, sexs)

	mockedBio := mockAndInsertBio(db, randCountryId, randCityId, sexs[0].Id, 1)

	assert.Equal(t, len(mockedBio), 1, "Mocking products failed")

	mockedGetSingleRequest := mockedGetSingleBioRequest(randUserId)

	fetchedBio, err := useCase.GetBio(ctx, "", mockedGetSingleRequest)
	assert.NotNil(t, fetchedBio)
	assertBioEquality(t, mockedBio, []Bio{*fetchedBio})
}

func createBioUseCase(db *gorm.DB, userId uint, withValidFunc bool) BioUseCaseInterface {
	validFunc := generateValidationFunction(withValidFunc)
	return NewBioUseCase(NewBioService(NewBioRepository(db), validFunc, validFunc), func(ctx context.Context, token string) (uint, error) {
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

func mockEditBioRequest(countryId, cityId, sexId uint, description string, born int64) *BioUpdateRequest {
	return &BioUpdateRequest{
		Description: description,
		Country:     countryId,
		City:        cityId,
		Sex:         sexId,
		Born:        born,
	}
}

func mockedGetSingleBioRequest(userId uint) *BioGetSingleRequest {
	return &BioGetSingleRequest{
		UserId: userId,
	}
}
