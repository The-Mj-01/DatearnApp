package uuid

import (
	"fmt"
	"github.com/gofrs/uuid"
	"strings"
)

// GenerateUUId and return it for user registration
func GenerateUUId(strings ...string) (string, error) {
	if strings != nil && len(strings) != 0 {
		return generateUUIDByStrings(strings...)
	}

	generatedUUID, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return generatedUUID.String(), nil
}

// generateUUIDByStrings by given strings
func generateUUIDByStrings(stringsSlice ...string) (string, error) {
	wholeString := strings.Join(stringsSlice, "_")
	generatedUUID, err := uuid.FromString(wholeString)
	if err != nil {
		return "", err
	}

	if generatedUUID.IsNil() {
		return "", fmt.Errorf("error on generating uuid:%s", "Invalid uuid type")
	}

	return generatedUUID.String(), nil
}
