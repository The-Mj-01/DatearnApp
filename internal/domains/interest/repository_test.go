package interest

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Interest{})
	return db, err
}

// createInterestRepo for testing purpose and return it
func createInterestRepo(db *gorm.DB) InterestRepositoryInterface {
	return NewInterestRepository(db)
}
