package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) Validate(model any) error {
	// Validate the struct
	if err := v.validate.Struct(model); err != nil {
		// Return validation errors as array
		validationErrors := make([]string, 0)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, getValidationMessage(err))
		}
		return errors.New(strings.Join(validationErrors, ", "))
	}

	return nil
}

// getValidationMessage returns a human-readable validation error message
func getValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s field is required", err.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", err.Field(), err.Param())
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", err.Field(), err.Param())
	case "numeric":
		return fmt.Sprintf("%s must be a valid number", err.Field())
	case "alpha":
		return fmt.Sprintf("%s must contain only alphabetic characters", err.Field())
	case "alphanum":
		return fmt.Sprintf("%s must contain only alphanumeric characters", err.Field())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}
