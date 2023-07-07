package image

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestImageService_GetAllImage(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createImageService(db)

	_, err = sv.GetAllImage(nil, nil, nil, nil, nil, 0)
	assert.Error(t, err, "Expected interest not found error")
	assert.ErrorIs(t, err, ImageNotFound, "Expected interest not found error")
}

func createImageService(db *gorm.DB) ImageServiceInterface {
	return NewImageService(NewImageRepository(db))
}
