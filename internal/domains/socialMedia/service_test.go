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
	assert.Error(t, err, "Expected countries not found error")
	assert.ErrorIs(t, err, SocialMediaNotFound, "Expected countries not found error")
}

func createSocialMediaService(db *gorm.DB) SocialMediaServiceInterface {
	return NewSocialMediaService(NewSocialMediaRepository(db))
}
