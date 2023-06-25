package socialMedia

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestSocialMediaUseCase_GetAllSocialMedia(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createSocialMediaUseCase(db, randUserId)

	mockedSocialMedia := mockAndInsertSocialMedia(db, 1)

	assert.Equal(t, len(mockedSocialMedia), 1, "Mocking products failed")

	mockedGetSocialMediaRequest := mockGetSocialMediaRequest(&mockedSocialMedia[0].Id, &mockedSocialMedia[0].Name, nil, 0)

	fetchedSocialMedia, err := useCase.GetAllSocialMedia(ctx, "", mockedGetSocialMediaRequest)
	assert.NotNil(t, fetchedSocialMedia)
	assertSocialMedia(t, mockedSocialMedia, *fetchedSocialMedia)

}

func createSocialMediaUseCase(db *gorm.DB, userId uint) SocialMediaUseCaseInterface {
	return NewSocialMediaUseCase(NewSocialMediaService(NewSocialMediaRepository(db)), func(ctx context.Context, token string) (uint, error) {
		return userId, nil
	})
}

func mockGetSocialMediaRequest(id *uint, name *string, limit *int, offset int) *SocialMediaGetRequest {
	return &SocialMediaGetRequest{
		Id:     id,
		Name:   name,
		Limit:  limit,
		Offset: offset,
	}
}
