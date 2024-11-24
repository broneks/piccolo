package service

import "github.com/go-playground/validator/v10"

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{validator.New(validator.WithRequiredStructEnabled())}
}

func (svc *Validator) Validate(i interface{}) error {
	return svc.validator.Struct(i)
}
