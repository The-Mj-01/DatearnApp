package api

import (
	"Datearn/internal/domains/user"
	"Datearn/internal/middleware/auth"
	"Datearn/pkg/IP"
	"Datearn/pkg/reqTokenHandler"
	"Datearn/pkg/validation"
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// userEchoHandler is the type which attaches rest api end points to domain functions
type userEchoHandler struct {
	uC user.UserUseCaseInterface
}

// AttachUserHandlerToUserDomain for working with rest Apis
func AttachUserHandlerToUserDomain(engine *echo.Echo, db *gorm.DB) {
	useCase := user.NewUserUseCase(user.NewService(user.NewRepository(db)))
	setupUserRoutes(engine, &userEchoHandler{
		uC: useCase,
	})
}

// setupUserRoutes which are accessible through http URI
func setupUserRoutes(engine *echo.Echo, handler *userEchoHandler) {
	userRouter := engine.Group("/users")

	userRouter.POST("", handler.Register)
	userRouter.POST("/login", handler.Login)

	userRouter.Use(auth.ValidateJWT)
	userRouter.PUT("/profile/username", handler.UpdateUserName)
	userRouter.PUT("/profile/pass", handler.UpdateUserPass)
}

// Register user in system
func (uH *userEchoHandler) Register(e echo.Context) error {
	request := new(user.UserRegisterRequest)

	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(request); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, user.ContextValueIpKey, IP.ExtractFromEcho(e))

	registerResp, err := uH.uC.Register(ctx, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(registerResp, nil))
}

// Login user into system
func (uH *userEchoHandler) Login(e echo.Context) error {
	request := new(user.UserLoginRequest)

	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(request); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, user.ContextValueIpKey, IP.ExtractFromEcho(e))

	loginResp, err := uH.uC.Login(ctx, request)
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(loginResp, nil))
}

// UpdateUserPass in system
func (uH *userEchoHandler) UpdateUserPass(e echo.Context) error {
	request := new(user.UpdateUserRequest)

	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(request); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	ctx := context.Background()

	updatedUser, err := uH.uC.UpdateUserPass(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(updatedUser, nil))
}

// UpdateUserName and return it
func (uH *userEchoHandler) UpdateUserName(e echo.Context) error {
	request := new(user.UpdateUserRequest)

	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(request); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	ctx := context.Background()

	updatedUser, err := uH.uC.UpdateUserName(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(updatedUser, nil))
}
