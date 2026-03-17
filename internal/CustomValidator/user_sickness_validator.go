package CustomValidator

import (
	"first-go-project/internal/model"

	"github.com/go-playground/validator/v10"
)

type UserSicknessValidator struct {
	Validate *validator.Validate
}

func NewUserSicknessValidator(validator *validator.Validate) *UserSicknessValidator {
	return &UserSicknessValidator{
		Validate: validator,
	}
}

func (c *UserSicknessValidator) ValidateCreateUserSickness(request *model.CreateUserSicknessRequest) error {
	return c.Validate.Struct(request)
}