package authorization

import (
	"errors"
	"reflect"
)

// wrongFieldError for cases when requested field doesn't exist in passed object
var wrongFieldError error = errors.New("requested field is not appropriate and did not found")

// SimpleFieldAuthorization is a simple authorizer which check if passed value is equal with requested field or not
func SimpleFieldAuthorization(obj any, requiredField string, expectedValue any, returningErr error) error {
	v := reflect.ValueOf(obj)

	field := v.FieldByName(requiredField)

	if !field.IsValid() {
		return wrongFieldError
	}

	fieldsVal := field.Interface()

	if fieldsVal != expectedValue {
		return returningErr
	}

	return nil
}
