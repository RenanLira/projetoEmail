package internalerrors

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {

	validate := validator.New()

	err := validate.Struct(obj)

	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]

	switch validationErr.Tag() {
	case "required":
		return errors.New("The field " + validationErr.Field() + " is required")
	case "email":
		return errors.New("The field " + validationErr.Field() + " is not a valid email")
	case "min":
		return errors.New("The field " + validationErr.Field() + " must have at least " + validationErr.Param() + " characters")
	case "max":
		return errors.New("The field " + validationErr.Field() + " must have at most " + validationErr.Param() + " characters")
	case "dive":
		return errors.New("The field " + validationErr.Field() + " must have at least one contact")
	default:
		return errors.New("The field " + validationErr.Field() + " is invalid")
	}
}
