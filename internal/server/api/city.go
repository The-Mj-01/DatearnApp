package api

import (
	"Datearn/internal/domains/location"
	"Datearn/internal/middleware/auth"
	"Datearn/pkg/reqTokenHandler"
	"Datearn/pkg/validation"
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// cityEchoHandler is the type which attaches rest api end points to domain functions
type cityEchoHandler struct {
	uC location.CityUseCaseInterface
}

// AttachCityToItsDomain for working with rest Apis
func AttachCityToItsDomain(engine *echo.Echo, db *gorm.DB) {
	uC := location.NewCityUseCase(location.NewCityService(location.NewCityRepository(db)), nil)
	setupCityRoutes(engine, &cityEchoHandler{
		uC: uC,
	})
}

// setupCityRoutes which are accessible through http URI
func setupCityRoutes(engine *echo.Echo, handler *cityEchoHandler) {
	router := engine.Group("/city")

	router.Use(auth.ValidateJWT)
	router.GET("/get_all", handler.GetAllCountries)

}

// GetAllLikes product for user
func (c *cityEchoHandler) GetAllCountries(e echo.Context) error {
	request := new(location.CityGetrequest)

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

	_, err = c.uC.GetAllCities(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}
