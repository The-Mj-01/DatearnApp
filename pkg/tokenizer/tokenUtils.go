package tokenizer

import (
	"Datearn/pkg/tokenizer/internal/JWT"
	"context"
	"fmt"
)

// DecodeTokenInfo and return it
func DecodeTokenInfo(headers map[string]string, ip string, ctx context.Context) (*JWT.Token, error) {
	authString, err := GetAuthString(headers)
	if err != nil {
		return nil, err
	}

	tokenizerHelper := CreateTokenizer(ctx)

	isActive := tokenizerHelper.IsActive(authString, ip)
	if !isActive {
		return nil, fmt.Errorf("%s", "Token is expired")
	}

	return tokenizerHelper.TokenInfo(authString)
}

// GetAuthString from passed header and return it
func GetAuthString(headers map[string]string) (string, error) {
	authString, exists := headers["Authorization"]

	if !exists {
		authString, exists = headers["authorization"]
	}

	if !exists {
		return "", fmt.Errorf("%s", "Token Doesnt exists")
	}

	return authString, nil
}
