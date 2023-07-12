package socialMedia

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestSocialMediaService_GetAllSocialMedia(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createSocialMediaService(db)

	_, err = sv.GetAllSocialMedia(nil, nil, nil, 0)
	assert.Error(t, err, "Expected social media not found error")
	assert.ErrorIs(t, err, SocialMediaNotFound, "Expected social media not found error")
}

func TestSocialMediaService_CreateSocialMedia(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createSocialMediaService(db)

	social := mockSocialMedia()

	createdSocialMedia, err := sv.CreateSocialMedia(social.Name)
	defer destructCreatedObjects(db, []SocialMedia{*createdSocialMedia})

	assert.NoError(t, err, "social media service bio creation failed")
	assert.Equal(t, social.Name, createdSocialMedia.Name)

}

func TestSocialMediaService_UpdateSocialMedia(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createSocialMediaService(db)

	oldSocialMedia := mockAndInsertSocialMedia(db, 1)
	defer destructCreatedObjects(db, oldSocialMedia)

	newSocialMedia := &SocialMedia{
		Id:   oldSocialMedia[0].Id,
		Name: "Twitter",
	}

	wrongSocialMedia := &SocialMedia{
		Name: "",
	}

	_, err = sv.UpdateSocialMedia(&newSocialMedia.Id, wrongSocialMedia.Name)
	assert.Error(t, err, "Social media service update functionality failed")
	assert.ErrorIs(t, err, NameNotFound, "Social media service update functionality failed")

	updatedSocialMedia, err := sv.UpdateSocialMedia(&newSocialMedia.Id, newSocialMedia.Name)

	assert.NoError(t, err, "Social media service update user failed")
	assert.Equal(t, newSocialMedia.Id, updatedSocialMedia.Id, "Social media service update bio failed")
	assert.Equal(t, newSocialMedia.Name, updatedSocialMedia.Name, "Social media service update bio failed")

}

func TestSocialMediaService_DeleteSocialMedia(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createSocialMediaService(db)
	social := mockAndInsertSocialMedia(db, 1)
	defer destructCreatedObjects(db, social)

	deletedUser, err := service.DeleteSocialMedia(&social[0].Id)

	assertSocialMedia(t, []SocialMedia{*deletedUser}, []SocialMedia{social[0]})

	_, err = service.DeleteSocialMedia(&social[0].Id)
	assert.Error(t, err, "Social service user creation failed")
	assert.ErrorIs(t, err, SocialMediaNotFound, "Social service user creation failed")

}

func createSocialMediaService(db *gorm.DB) SocialMediaServiceInterface {
	return NewSocialMediaService(NewSocialMediaRepository(db))
}
