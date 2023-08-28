package helper

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) ValidateStruct(s interface{}) error {
	err := v.validator.Struct(s)

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range castedObject {
			switch fieldError.Tag() {
			case "required":
				err = NewServiceError(http.StatusBadRequest, fmt.Sprintf("%s is required", fieldError.Field()))
			}
		}
	}

	return err
}
