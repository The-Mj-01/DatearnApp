package location

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestCountryUseCase_GetAllCountries(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createCountryUseCase(db, randUserId)

	mockedCountry := mockAndInsertCountry(db, 10)
	defer destructCreatedObjects(db, mockedCountry)
	assert.Equal(t, len(mockedCountry), 10, "Mocking products failed")

	mockedGetAllCountryRequest := mockGetAllCountryRequest(&mockedCountry[0].Name, nil, 0)

	fetchedCountry, err := useCase.GetAllCountries(ctx, "", mockedGetAllCountryRequest)
	assert.NoError(t, err, "Fetching Countries from db failed")
	assertCountries(t, mockedCountry, *fetchedCountry)
}

func TestCityUseCase_GetAllCities(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createCityUseCase(db, randUserId)

	mockedCity := mockAndInsertCity(db, 10)
	defer destructCreatedObjects(db, mockedCity)
	assert.Equal(t, len(mockedCity), 10, "Mocking products failed")

	mockedGetAllCityRequest := mockGetAllCityRequest(&mockedCity[0].Name, nil, 0)

	fetchedCity, err := useCase.GetAllCities(ctx, "", mockedGetAllCityRequest)
	assert.NoError(t, err, "Fetching Cities from db failed")
	assertCities(t, mockedCity, *fetchedCity)

}

func createCountryUseCase(db *gorm.DB, userId uint) CountryUseCaseInterface {
	return NewCountryUseCase(NewCountryService(NewCountryRepository(db)), func(ctx context.Context, token string) (uint, error) {
		return userId, nil
	})
}

func createCityUseCase(db *gorm.DB, userId uint) CityUseCaseInterface {
	return NewCityUseCase(NewCityService(NewCityRepository(db)), func(ctx context.Context, token string) (uint, error) {
		return userId, nil
	})
}

func mockGetAllCountryRequest(name *string, limit *int, offset int) *CountryGetrequest {
	return &CountryGetrequest{
		Name:   name,
		Limit:  limit,
		Offset: offset,
	}
}

func mockGetAllCityRequest(name *string, limit *int, offset int) *CityGetrequest {
	return &CityGetrequest{
		Name:   name,
		Limit:  limit,
		Offset: offset,
	}
}
