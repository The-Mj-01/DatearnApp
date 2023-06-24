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

func createSocialMediaService(db *gorm.DB) SocialMediaServiceInterface {
	return NewSocialMediaService(NewSocialMediaRepository(db))
}
