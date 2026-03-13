package CustomValidator

import (
	"first-go-project/internal/model"

	"github.com/go-playground/validator/v10"
)

type UserValidator struct {
	Validate *validator.Validate
}

func NewUserValidator(validator *validator.Validate) *UserValidator {
	return &UserValidator{
		Validate: validator,
	}
}

func (c *UserValidator) ValidateCreateUser(request *model.CreateUserRequest) error {
	return c.Validate.Struct(request)
}

func (c *UserValidator) ValidateUpdateUser(request *model.UpdateUserRequest) error {
	return c.Validate.Struct(request)
}