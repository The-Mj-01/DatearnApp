package interest

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestInterestService_GetAllInterest(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	sv := createInterestService(db)

	_, err = sv.GetAllInterest(nil, nil, nil, 0)
	assert.Error(t, err, "Expected interest not found error")
	assert.ErrorIs(t, err, InterestNotFound, "Expected interest not found error")
}

func createInterestService(db *gorm.DB) InterestServiceInterface {
	return NewInterestService(NewInterestRepository(db))
}
