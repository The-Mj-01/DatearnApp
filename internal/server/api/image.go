package api

import (
	"Datearn/internal/domains/image"
	"Datearn/pkg/reqTokenHandler"
	"Datearn/pkg/validation"
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// imageEchoHandler is the type which attaches rest api end points to domain functions
type imageEchoHandler struct {
	uC image.ImageUseCaseInterface
}

// AttachImageToItsDomain for working with rest Apis
func AttachImageToItsDomain(engine *echo.Echo, db *gorm.DB) {
	uC := image.NewImageUseCase(image.NewImageService(image.NewImageRepository(db)), nil)
	setupImageRoutes(engine, &imageEchoHandler{
		uC: uC,
	})
}

// setupImageRoutes which are accessible through http URI
func setupImageRoutes(engine *echo.Echo, handler *imageEchoHandler) {
	router := engine.Group("image")
	router.GET("/get_all", handler.GetAllImage)
	router.POST("/create", handler.CreateImage)
	router.PUT("/update/:id", handler.UpdateImage)
	router.DELETE("/delete/:id", handler.DeleteImage)
}

// GetAllImage product for user
func (s *imageEchoHandler) GetAllImage(e echo.Context) error {
	request := new(image.ImageGetRequest)

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

	_, err = s.uC.GetAllImage(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// CreateImage product for user
func (s *imageEchoHandler) CreateImage(e echo.Context) error {
	request := new(image.ImageCreateRequest)

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

	_, err = s.uC.CreateImage(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// UpdateImage product for user
func (s *imageEchoHandler) UpdateImage(e echo.Context) error {
	request := new(image.ImageUpdateRequest)

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

	_, err = s.uC.UpdateImage(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// DeleteImage product for user
func (s *imageEchoHandler) DeleteImage(e echo.Context) error {
	request := new(image.ImageDeleteRequest)

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

	_, err = s.uC.DeleteImage(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}
