package api

import (
	"Datearn/internal/domains/interest"
	"Datearn/pkg/reqTokenHandler"
	"Datearn/pkg/validation"
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// interestEchoHandler is the type which attaches rest api end points to domain functions
type interestEchoHandler struct {
	uC interest.InterestUseCaseInterface
}

// AttachInterestToItsDomain for working with rest Apis
func AttachInterestToItsDomain(engine *echo.Echo, db *gorm.DB) {
	uC := interest.NewInterestUseCase(interest.NewInterestService(interest.NewInterestRepository(db)), nil)
	setupInterestRoutes(engine, &interestEchoHandler{
		uC: uC,
	})
}

// setupInterestRoutes which are accessible through http URI
func setupInterestRoutes(engine *echo.Echo, handler *interestEchoHandler) {
	router := engine.Group("interest")
	router.POST("/get_all", handler.GetAllInterest)
	router.POST("/create", handler.CreateInterest)
	router.POST("/update", handler.UpdateInterest)
	router.POST("/delete", handler.DeleteInterest)
}

// GetAllInterest product for user
func (s *interestEchoHandler) GetAllInterest(e echo.Context) error {
	request := new(interest.InterestGetRequest)

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

	_, err = s.uC.GetAllInterest(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// CreateInterest product for user
func (s *interestEchoHandler) CreateInterest(e echo.Context) error {
	request := new(interest.InterestCreateRequest)

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

	_, err = s.uC.CreateInterest(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// UpdateInterest product for user
func (s *interestEchoHandler) UpdateInterest(e echo.Context) error {
	request := new(interest.InterestUpdateRequest)

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

	_, err = s.uC.UpdateInterest(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// DeleteInterest product for user
func (s *interestEchoHandler) DeleteInterest(e echo.Context) error {
	request := new(interest.InterestDeleteRequest)

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

	_, err = s.uC.DeleteInterest(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}
