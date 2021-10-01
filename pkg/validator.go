package pkg

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	ValidationUtil struct {
		validator *validator.Validate
	}
)

func NewValidationUtil() echo.Validator {
	return &ValidationUtil{validator: validator.New()}
}

func (v *ValidationUtil) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

type ValidationError struct {
	Key     string      `json:"key"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}

func ParseValidationError(err error) []ValidationError {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return nil
	}

	var validationErrors []ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, ValidationError{
			Key:     err.StructField(),
			Value:   err.Value(),
			Message: err.Tag(),
		})
	}
	return validationErrors
}
