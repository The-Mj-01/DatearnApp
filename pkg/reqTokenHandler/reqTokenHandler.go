package reqTokenHandler

import (
	"errors"
	"net/http"
	"strings"
)

// bearerTokenKeyIndicator for indicating passed token from client
const bearerTokenKeyIndicator string = "Authorization"

var noBearerProvided error = errors.New("no bearer token passed by client")

// ExtractBearerToken and return it if any exists
func ExtractBearerToken(req *http.Request) (string, error) {
	bearerToken := req.Header.Get(bearerTokenKeyIndicator)

	if bearerToken != "" {
		goto EndLine
	}

	bearerToken = req.Header.Get(strings.ToLower(bearerTokenKeyIndicator))
	if bearerToken == "" {
		return "", noBearerProvided
	}

EndLine:
	return bearerToken, nil
}
