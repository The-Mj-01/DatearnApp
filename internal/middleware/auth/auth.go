package auth

import (
	"Datearn/pkg/IP"
	"Datearn/pkg/reqTokenHandler"
	"Datearn/pkg/tokenizer"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

// unAuthorizedMsg defines a message for authorization
var unAuthorizedMsg map[string]string = map[string]string{
	"data":    "",
	"message": "You are not authorized",
}

// ValidateJWT for user
func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bearerToken := c.Request().Header.Get("Authorization")
		if bearerToken == "" {
			return c.JSON(http.StatusUnauthorized, unAuthorizedMsg)
		}

		bearerToken = reqTokenHandler.ClearTokenString(bearerToken)

		ip := IP.ExtractFromEcho(c)
		isAllowed := handleJwtValidation(bearerToken, ip)

		if !isAllowed {
			return c.JSON(http.StatusUnauthorized, unAuthorizedMsg)
		}

		return next(c)
	}
}

// handleJwtValidation handles JWT verification functionality
func handleJwtValidation(token, ip string) bool {
	ctx := context.Background()
	tokenHandler := tokenizer.CreateTokenizer(ctx)
	return tokenHandler.IsActive(token, ip)
}
