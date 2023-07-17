package api

import (
	"Datearn/internal/domains/location"
	"Datearn/pkg/reqTokenHandler"
	"Datearn/pkg/validation"
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// countryEchoHandler is the type which attaches rest api end points to domain functions
type countryEchoHandler struct {
	uC location.CountryUseCaseInterface
}

// AttachCountryToItsDomain for working with rest Apis
func AttachCountryToItsDomain(engine *echo.Echo, db *gorm.DB) {
	uC := location.NewCountryUseCase(location.NewCountryService(location.NewCountryRepository(db)), nil)
	setupCountryRoutes(engine, &countryEchoHandler{
		uC: uC,
	})
}

// setupCountryRoutes which are accessible through http URI
func setupCountryRoutes(engine *echo.Echo, handler *countryEchoHandler) {
	router := engine.Group("swipe")
	router.GET("/get_all", handler.GetAllCountries)

}

// GetAllLikes product for user
func (c *countryEchoHandler) GetAllCountries(e echo.Context) error {
	request := new(location.CountryGetrequest)

	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(request); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	ctx := context.Background()

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	_, err = c.uC.GetAllCountries(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}
