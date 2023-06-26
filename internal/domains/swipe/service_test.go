package swipe

import (
	"gorm.io/gorm"
)

func createSwipeService(db *gorm.DB) SwipeServiceInterface {
	return NewSwipeService(NewSwipeRepository(db))
}
