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

func TestSocialMediaUseCase_CreateSocialMedia(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createSocialMediaUseCase(db, randUserId)

	social := mockSocialMedia()

	mockedRequest := mockWriteSocialMediaRequest(social.Name)
	result, err := useCase.CreateSocialMedia(ctx, "", mockedRequest)
	assert.NoError(t, err, "Social media creation failed in address use-case")
	assert.Equal(t, result.Name, mockedRequest.Name, "Social media creation failed in bio use-case")

}

func TestSocialMediaUseCase_UpdateSocialMedia(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createSocialMediaUseCase(db, randUserId)

	oldSocialMedia := mockAndInsertSocialMedia(db, 1)
	defer destructCreatedObjects(db, oldSocialMedia)

	newName := "Twitter"
	mockedEditRequest := mockEditSocialMediaRequest(&oldSocialMedia[0].Id, &newName)
	editedSocialMedia, err := useCase.UpdateSocialMedia(ctx, "", mockedEditRequest)
	assert.NoError(t, err, "Social media use-case update functionality failed")

	assert.Equal(t, *mockedEditRequest.Name, editedSocialMedia.Name, "Social media use-case update functionality failed")
}

func TestSocialMediaUseCase_DeleteSocialMedia(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.Background()
	randUserId := uint(1)
	useCase := createSocialMediaUseCase(db, randUserId)

	mockedSocialMedia := mockAndInsertSocialMedia(db, 1)
	defer destructCreatedObjects(db, mockedSocialMedia)

	mockedDeleteRequest := mockDeleteSocialMediaRequest(&mockedSocialMedia[0].Id)

	deletedSocialMedia, err := useCase.DeleteSocialMedia(ctx, "", mockedDeleteRequest)
	assert.NoError(t, err, "Deleting user name failed")
	//fmt.Println(user, mockedUser)
	assertSocialMedia(t, mockedSocialMedia, []SocialMedia{*deletedSocialMedia})

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

func mockWriteSocialMediaRequest(name string) *SocialMediaCreateRequest {
	return &SocialMediaCreateRequest{
		Name: name,
	}
}

func mockEditSocialMediaRequest(id *uint, name *string) *SocialMediaUpdateRequest {
	return &SocialMediaUpdateRequest{
		Id:   id,
		Name: name,
	}
}

func mockDeleteSocialMediaRequest(id *uint) *SocialMediaDeleteRequest {
	return &SocialMediaDeleteRequest{
		Id: id,
	}
}
