package http

import (
	"first-go-project/internal/model"
	"first-go-project/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AddressController struct {
	Log     *logrus.Logger
	UseCase *usecase.AddressUseCase
}

func NewAddressController(useCase *usecase.AddressUseCase, logger *logrus.Logger) *AddressController {
	return &AddressController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *AddressController) UpdateAddress(ctx *fiber.Ctx) error {
	request := new(model.UpdateAddressRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.UpdateAddress(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to update address : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UpdateAddressResponse]{Data: response})
}

func (c *AddressController) DeleteAddress(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	response, err := c.UseCase.DeleteAddress(ctx.UserContext(), idStr)
	if err != nil {
		c.Log.Warnf("Failed to delete address : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.DeleteAddressResponse]{Data: response})
}