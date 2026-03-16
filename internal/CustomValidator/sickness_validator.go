package CustomValidator

import (
	"first-go-project/internal/model"

	"github.com/go-playground/validator/v10"
)

type SicknessValidator struct {
	Validate *validator.Validate
}

func NewSicknessValidator(validator *validator.Validate) *SicknessValidator {
	return &SicknessValidator{
		Validate: validator,
	}
}


func (c *SicknessValidator) ValidateCreateSickness(request *model.CreateSicknessRequest) error {
	return c.Validate.Struct(request)
}

func (c *SicknessValidator) ValidateUpdateSickness(request *model.UpdateSicknessRequest) error {
	return c.Validate.Struct(request)
}