package api

import (
	"Datearn/pkg/validation"
	"encoding/json"
	"fmt"
)

// generateResponse that is suitable for api and return it
func generateResponse(data any, message any) map[string]interface{} {
	return map[string]interface{}{
		"data":    data,
		"message": convertMessage(message),
	}
}

// convertMessage to string and return it
func convertMessage(message any) string {
	var convertedMsg string

	switch val := message.(type) {
	case error:
		convertedMsg = val.Error()
	case string:
		convertedMsg = val
	case []*validation.ValidationError:
		convertedMsg = convertValidationErrors(val)
	default:
		convertedMsg = ""
	}

	return convertedMsg
}

// convertValidationErrors into a single string
func convertValidationErrors(errs []*validation.ValidationError) string {
	marshaledData, err := json.Marshal(errs)
	if err != nil {
		return ""
	}

	return string(marshaledData)
}

// generateVerificationMessage on success request
func generateVerificationMessage(id int) string {
	return fmt.Sprintf("Verification request submitted successfully with id: %d", id)
}
