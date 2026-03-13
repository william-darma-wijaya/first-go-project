package CustomValidator

import (
	"first-go-project/internal/model"

	"github.com/go-playground/validator/v10"
)

type AddressValidator struct {
	Validate *validator.Validate
}

func NewAddressValidator(validator *validator.Validate) *AddressValidator {
	return &AddressValidator{
		Validate: validator,
	}
}

func (c *AddressValidator) ValidateUpdateAddress(request *model.UpdateAddressRequest) error {
	return c.Validate.Struct(request)
}