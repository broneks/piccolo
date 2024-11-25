package validatorservice

func (svc *ValidatorService) Validate(s any) error {
	return svc.validator.Struct(s)
}
