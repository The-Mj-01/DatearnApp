package JWT

import (
	"Datearn/pkg/advancedError"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

const (
	TOKEN_EXP_TIME time.Duration = time.Minute * 60
	REF_TOKEN_EXP  time.Duration = time.Minute * 10080
	REF_KEY_STR    string        = "REF"
	JWT_KEY_ENVKEY string        = "JWT_KEY"
)

// jwtKey for token Creation
var jwtKey = []byte(os.Getenv(JWT_KEY_ENVKEY))

// Token implementation for tokenizer
type Token struct {
	Ctx  context.Context `json:"-"`
	UUID string          `json:"uuid,string"`
	Name string          `json:"name,string"`
	//Exp  int64           `json:"exp,int64"`
	Ip string `json:"ip,string"`
	jwt.StandardClaims
}

// New jwt token and returns it
func (t *Token) New(uuid, ip string) (map[string]string, error) {
	t.UUID = uuid
	t.Ip = ip
	t.ExpiresAt = t.generateExpTime("")

	tokenString, err := t.generateTokenStr()
	if err != nil {
		return nil, advancedError.New(err, "Cannot create new token string")
	}

	//Generate refresh token
	t.ExpiresAt = t.generateExpTime(REF_KEY_STR)

	refTokenString, err := t.generateTokenStr()

	if err != nil {
		return nil, advancedError.New(err, "Cannot create new refresh token string")
	}

	return t.generateTksResp(tokenString, refTokenString), nil
}

// Refresh expired token and return it
func (t *Token) Refresh(token, ip string) (string, error) {
	if !t.IsActive(token, ip) {
		return "", fmt.Errorf("%s", "Token is not active")
	}

	err := t.Decrypt(token)
	if err != nil {
		return "", err
	}

	t.ExpiresAt = t.generateExpTime("")

	tokenString, err := t.generateTokenStr()
	if err != nil {
		return "", advancedError.New(err, "Cannot generate new token with refresh token")
	}

	fmt.Println(tokenString)

	return tokenString, nil
}

// IsActive checks if it's active or not
func (t *Token) IsActive(token, ip string) bool {
	err := t.Decrypt(token)

	if err != nil {
		return false
	}

	if ip != t.Ip {
		return false
	}

	if tokenIsExpired(t.ExpiresAt) {
		return false
	}

	return true
}

// TokenInfo decoded token and returns it
func (t *Token) TokenInfo(token string) (*Token, error) {
	err := t.Decrypt(token)
	if err != nil {
		return nil, advancedError.New(err, "Cannot return token info")
	}

	return t, nil
}

// Decrypt Sent Token
func (t *Token) Decrypt(sentToken string) error {
	tknParseInfo, err := jwt.ParseWithClaims(sentToken, t, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return advancedError.New(err, "Cannot Decrypt token")
	}

	if !tknParseInfo.Valid {
		return advancedError.New(err, "Decoded token is invalid")
	}

	return nil
}

// generateExpTime for token
func (t *Token) generateExpTime(tkType string) int64 {
	if tkType == REF_KEY_STR {
		return time.Now().Unix() + int64(REF_TOKEN_EXP.Seconds())
	}

	return time.Now().Unix() + int64(TOKEN_EXP_TIME.Seconds())
}

// tokenIsExpired for checking if token is expired or not
func tokenIsExpired(tkExpire int64) bool {
	if time.Now().Unix() > tkExpire {
		return true
	}
	return false
}

func (t *Token) generateTokenStr() (string, error) {
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, t)

	tokenString, err := tokenClaim.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// generateTksResp for response generation
func (t *Token) generateTksResp(token string, refToken string) map[string]string {
	resp := make(map[string]string)
	resp["token"] = token
	resp["refresh_token"] = refToken
	return resp
}
