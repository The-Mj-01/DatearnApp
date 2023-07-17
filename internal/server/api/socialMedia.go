package api

import (
	"Datearn/internal/domains/socialMedia"
	"Datearn/pkg/reqTokenHandler"
	"Datearn/pkg/validation"
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// socialMediaEchoHandler is the type which attaches rest api end points to domain functions
type socialMediaEchoHandler struct {
	uC socialMedia.SocialMediaUseCaseInterface
}

// AttachSocialMediaToItsDomain for working with rest Apis
func AttachSocialMediaToItsDomain(engine *echo.Echo, db *gorm.DB) {
	uC := socialMedia.NewSocialMediaUseCase(socialMedia.NewSocialMediaService(socialMedia.NewSocialMediaRepository(db)), nil)
	setupSocialMediaRoutes(engine, &socialMediaEchoHandler{
		uC: uC,
	})
}

// setupSocialMediaRoutes which are accessible through http URI
func setupSocialMediaRoutes(engine *echo.Echo, handler *socialMediaEchoHandler) {
	router := engine.Group("socialMedia")
	router.GET("/get_all", handler.GetAllSocialMedia)
	router.POST("/create", handler.CreateSocialMedia)
	router.PUT("/update/:id", handler.UpdateSocialMedia)
	router.DELETE("/delete/:id", handler.DeleteSocialMedia)
}

// GetAllSocialMedia product for user
func (s *socialMediaEchoHandler) GetAllSocialMedia(e echo.Context) error {
	request := new(socialMedia.SocialMediaGetRequest)

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

	_, err = s.uC.GetAllSocialMedia(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// CreateSocialMedia product for user
func (s *socialMediaEchoHandler) CreateSocialMedia(e echo.Context) error {
	request := new(socialMedia.SocialMediaCreateRequest)

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

	_, err = s.uC.CreateSocialMedia(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// UpdateSocialMedia product for user
func (s *socialMediaEchoHandler) UpdateSocialMedia(e echo.Context) error {
	request := new(socialMedia.SocialMediaUpdateRequest)

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

	_, err = s.uC.UpdateSocialMedia(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// DeleteSocialMedia product for user
func (s *socialMediaEchoHandler) DeleteSocialMedia(e echo.Context) error {
	request := new(socialMedia.SocialMediaDeleteRequest)

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

	_, err = s.uC.DeleteSocialMedia(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}
