package tokenizer

import (
	"Datearn/pkg/tokenizer/internal/JWT"
	"context"
)

// Tokenizer defines authentication operations checker
type Tokenizer interface {
	//New Create and return new Token
	New(uuid, ip string) (map[string]string, error)
	//Refresh expired token and return it
	Refresh(token, ip string) (string, error)
	//IsActive checks whether token is active or not
	IsActive(token, ip string) bool
	//TokenInfo decodes token and returns it
	TokenInfo(token string) (*JWT.Token, error)
	//Decrypt exiting token
	Decrypt(token string) error
}

// CreateTokenizer is tokenizer factory
func CreateTokenizer(ctx context.Context) Tokenizer {
	return &JWT.Token{Ctx: ctx}
}
