package api

import (
	"Datearn/internal/domains/swipe"
	"Datearn/internal/middleware/auth"
	"Datearn/pkg/reqTokenHandler"
	"Datearn/pkg/validation"
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// operationSuccess is a message for successful operations
const operationSuccess string = "Operation was successful"

// swipeEchoHandler is the type which attaches rest api end points to domain functions
type swipeEchoHandler struct {
	uC swipe.SwipeUseCaseInterface
}

// AttachSwipeToItsDomain for working with rest Apis
func AttachSwipeToItsDomain(engine *echo.Echo, db *gorm.DB) {
	uC := swipe.NewSwipeUseCase(swipe.NewSwipeService(swipe.NewSwipeRepository(db)), nil)
	setupSwipeRoutes(engine, &swipeEchoHandler{
		uC: uC,
	})
}

// setupSwipeRoutes which are accessible through http URI
func setupSwipeRoutes(engine *echo.Echo, handler *swipeEchoHandler) {
	router := engine.Group("swipe")
	router.Use(auth.ValidateJWT)
	router.POST("/like", handler.Like)
	router.POST("/disable_like", handler.DisableLike)
	router.POST("/dislike", handler.Dislike)
	router.POST("/disable_dislike", handler.DisableDisLike)
	router.GET("/get_all", handler.GetAllLikes)
}

// Like product for user
func (s *swipeEchoHandler) Like(e echo.Context) error {
	request := new(swipe.LikeRequest)

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

	_, err = s.uC.Like(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// DisableLike product for user
func (s *swipeEchoHandler) DisableLike(e echo.Context) error {
	request := new(swipe.LikeRequest)

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

	_, err = s.uC.DisableLike(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// Dislike product for user
func (s *swipeEchoHandler) Dislike(e echo.Context) error {
	request := new(swipe.DisLikeRequest)

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

	_, err = s.uC.DisLike(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// DisableDisLike product for user
func (s *swipeEchoHandler) DisableDisLike(e echo.Context) error {
	request := new(swipe.DisLikeRequest)

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

	_, err = s.uC.DisableDisLike(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// GetAllLikes product for user
func (s *swipeEchoHandler) GetAllLikes(e echo.Context) error {
	request := new(swipe.GetLikeRequest)

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

	_, err = s.uC.GetAllLikes(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}
