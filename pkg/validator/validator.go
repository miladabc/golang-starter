package validator

import (
	validator "github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
)

type EchoValidator struct {
	validator *validator.Validate
}

func NewEchoValidator() echo.Validator {
	return &EchoValidator{validator: validator.New()}
}

func (v *EchoValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
