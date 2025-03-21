package utils

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type EchoValidator struct {
	validator *validator.Validate
}

func (v *EchoValidator) Validate(i interface{}) error {
	result := v.validator.Struct(i)
	if result != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error())
	}
	return result
}

func (v *EchoValidator) SetValidator(newValidator *validator.Validate) {
	v.validator = newValidator
}

func NewEchoValidator() *EchoValidator {
	return &EchoValidator{
		validator: validator.New(),
	}
}
