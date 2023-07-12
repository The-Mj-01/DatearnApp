package JWT

import (
	"context"
	"fmt"
	"testing"
)

const (
	tokeKeyName string = "token"
	refTokenKey string = "refresh_token"
	ipString    string = "192.168.1.1"
)

// TestToken_New for testing creating jwt token
func TestToken_New(t *testing.T) {
	_, testTk := createToken(t)

	_, exists := testTk[tokeKeyName]
	if !exists {
		t.Error("Test getting JWT token but it doesnt exists")
	}

	_, exists = testTk[refTokenKey]
	if !exists {
		t.Error("Test getting JWT refresh token but it doesnt exists")
	}
}

// TestToken_Decrypt for decryption operation
func TestToken_Decrypt(t *testing.T) {
	_, testTk := createToken(t)

	decryptor := Token{
		Ctx: context.Background(),
	}

	err := decryptor.Decrypt(testTk[tokeKeyName])
	if err != nil {
		t.Errorf("%s %s", "Testing JWT decryption, expected decryption but got", err)
	}

	if decryptor.UUID == "" || decryptor.Ip == "" {
		t.Error("Testing JWT decryption, expected decryption but got null struct field")
	}

	if decryptor.Ip != ipString || decryptor.UUID != "iamMj" {
		t.Error("Testing JWT decryption but decoded params are not equal")
	}
}

// TestToken_IsActive Checks if token is active or not
func TestToken_IsActive(t *testing.T) {
	tkObj, testTk := createToken(t)

	if !tkObj.IsActive(testTk[tokeKeyName], ipString) {
		t.Error("Newly generated token is expected to be active but it's not")
	}

	if !tkObj.IsActive(testTk[refTokenKey], ipString) {
		t.Error("Newly generated refresh token is expected to be active but it's not")
	}
}

// TestToken_Refresh
func TestToken_Refresh(t *testing.T) {
	tkObj, testTk := createToken(t)
	fmt.Println(testTk[refTokenKey])
	refTkString, err := tkObj.Refresh(testTk[refTokenKey], ipString)
	if err != nil {
		t.Errorf("%s: %s", "Expected To do refresh operation but got", err)
	}

	if refTkString != testTk[tokeKeyName] {
		t.Error("Expected That newly generated token not be the same with old token")
	}
}

// createToken and return it
func createToken(t *testing.T) (Token, map[string]string) {
	jwt := Token{
		Ctx: context.Background(),
	}

	testTk, err := jwt.New("iamMj", ipString)
	if err != nil {
		t.Errorf("%s %s", "Testing JWT creation, expected jwt but got", err)
	}
	return jwt, testTk
}
