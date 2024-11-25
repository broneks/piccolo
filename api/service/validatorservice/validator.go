package validatorservice

import "github.com/go-playground/validator/v10"

type ValidatorService struct {
	validator *validator.Validate
}

func New() *ValidatorService {
	return &ValidatorService{validator.New(validator.WithRequiredStructEnabled())}
}
