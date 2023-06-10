package validation

import "github.com/go-playground/validator"

// ValidationError represents validation error
type ValidationError struct {
	FailedField string
	Tag         string
	Value       string
}

// Validate and return error is any data is wrong
func Validate(data interface{}) []*ValidationError {
	var errors []*ValidationError
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
