package http

import (
	"Datearn/internal/server/api"
	"Datearn/pkg/advancedError"
	"Datearn/pkg/database"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"net/http"
	"os"
)

// DEFAULT_PORT defines default application port
const DEFAULT_PORT string = "8000"

// StartHttpServer for running
func StartHttpServer() {
	dbConn := connectToDb()
	router := makeNewApp()

	registerServices(dbConn, router)

	port := appPort()
	if err := router.Start(port); err != nil {
		panic(err)
	}
}

// connectToDb initializes database connection
func connectToDb() *gorm.DB {
	conn, err := database.Conn()
	if err != nil {
		panic(advancedError.New(err, "Starting database connection failed"))
	}
	return conn
}

// makeNewApp and return it
func makeNewApp() *echo.Echo {
	return echo.New()
}

// registerServices in main app
func registerServices(conn *gorm.DB, router *echo.Echo) {
	appendRequiredMiddlewares(router)
	api.AttachBioToItsDomain(router, conn)
	api.AttachImageToItsDomain(router, conn)
	api.AttachInterestToItsDomain(router, conn)
	api.AttachCountryToItsDomain(router, conn)
	api.AttachCityToItsDomain(router, conn)
	api.AttachSocialMediaToItsDomain(router, conn)
	api.AttachSwipeToItsDomain(router, conn)
	api.AttachUserHandlerToUserDomain(router, conn)
}

// appendRequiredMiddlewares to http router function
func appendRequiredMiddlewares(router *echo.Echo) {
	router.Use(middleware.Logger())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	//Todo disable strict routing
}

// appPort gets and returns application port
func appPort() string {
	port := os.Getenv("APPLICATION_PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	return fmt.Sprintf(":%s", port)
}
