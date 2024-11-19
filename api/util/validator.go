package util

import "github.com/go-playground/validator/v10"

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{validator.New(validator.WithRequiredStructEnabled())}
}

func (va *Validator) Validate(i interface{}) error {
	return va.validator.Struct(i)
}
