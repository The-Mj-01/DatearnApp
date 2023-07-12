package bioHandler

import (
	"Datearn/internal/domains/location"
	"Datearn/pkg/database"
	"errors"
	"gorm.io/gorm"
	"sync"
)

var (
	CountryIsInvalid error = errors.New("selected country is invalid")
	CityIsInvalid    error = errors.New("selected city is invalid")
)

var CountryRepo location.CountryRepositoryInterface
var CityRepo location.CityRepositoryInterface
var once sync.Once

// init initializes user handler
func initFunc() {
	db := connectToDb()
	CountryRepo = location.NewCountryRepository(db)
	CityRepo = location.NewCityRepository(db)

}

// connectToDb makes database connection
func connectToDb() *gorm.DB {
	conn, err := database.Conn()
	if err != nil {
		panic(err)
	}
	return conn
}

// CountryIsValid checks whether given basket exists or not
func CountryIsValid(countryId uint) error {
	once.Do(initFunc)
	exists := CountryRepo.CountryExists(countryId)
	if !exists {
		return CountryIsInvalid
	}
	return nil
}

// CityIsValid checks whether given basket exists or not
func CityIsValid(cityId uint) error {
	once.Do(initFunc)
	exists := CityRepo.CityExists(cityId)
	if !exists {
		return CityIsInvalid
	}
	return nil
}
