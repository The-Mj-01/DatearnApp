package api

import (
	"Datearn/internal/domains/bio"
	"Datearn/internal/middleware/auth"
	"Datearn/pkg/reqTokenHandler"
	"Datearn/pkg/validation"
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// Todo: write system test for APIs
// bioEchoHandler is the type which attaches rest api end points to domain functions
type bioEchoHandler struct {
	uC bio.BioUseCaseInterface
}

// AttachBioToItsDomain for working with rest Apis
func AttachBioToItsDomain(engine *echo.Echo, db *gorm.DB) {
	uC := bio.NewBioUseCase(bio.NewBioService(bio.NewBioRepository(db), nil, nil), nil)
	setupBioRoutes(engine, &bioEchoHandler{
		uC: uC,
	})
}

// setupBioRoutes which are accessible through http URI
func setupBioRoutes(engine *echo.Echo, handler *bioEchoHandler) {
	router := engine.Group("bio")
	router.Use(auth.ValidateJWT)
	router.GET("/get/:id", handler.GetBio)
	router.POST("/create", handler.WriteBio)
	router.PUT("/update", handler.UpdateBio)
}

// WriteBio product for bio
func (b *bioEchoHandler) WriteBio(e echo.Context) error {
	request := new(bio.BioCreateRequest)

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

	_, err = b.uC.WriteBio(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// GetBio product for bio
func (b *bioEchoHandler) GetBio(e echo.Context) error {
	request := new(bio.BioGetSingleRequest)

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

	_, err = b.uC.GetBio(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// UpdateBio product for bio
func (b *bioEchoHandler) UpdateBio(e echo.Context) error {
	request := new(bio.BioUpdateRequest)

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

	_, err = b.uC.UpdateBio(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}
