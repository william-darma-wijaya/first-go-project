package http

import (
	"first-go-project/internal/model"
	"first-go-project/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SicknessController struct {
	Log     *logrus.Logger
	UseCase *usecase.SicknessUsecase
}

func NewSicknessController(useCase *usecase.SicknessUsecase, logger *logrus.Logger) *SicknessController {
	return &SicknessController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *SicknessController) CreateNewSickness(ctx *fiber.Ctx) error {
	request := new(model.CreateSicknessRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse body request: %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.CreateSickness(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create new sickness: %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CreateSicknessResponse]{Data: response})
}

func (c *SicknessController) GetSicknessById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	response, err := c.UseCase.FindById(ctx.UserContext(), idStr)
	if err != nil {
		c.Log.Warnf("Failed to get sickness by Id: %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.GetSicknessByIdResponse]{Data: response})
}

func (c *SicknessController) UpdateSickness(ctx *fiber.Ctx) error {
	request := new(model.UpdateSicknessRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse body request: %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.UpdateSickness(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to update user: %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UpdateSicknessResponse]{Data: response})
}

func (c *SicknessController) DeleteSickness(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	response, err := c.UseCase.DeleteSickness(ctx.UserContext(), idStr)
	if err != nil {
		c.Log.Warnf("Failed to delete sickness: %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.DeleteSicknessResponse]{Data: response})
}

func (c *SicknessController) FindSicknesses(ctx *fiber.Ctx) error {
	name := ctx.Query("name")

	response, err := c.UseCase.FindSicknessesByName(ctx.UserContext(), name)
	if err != nil {
		c.Log.Warnf("Failed to retrieve sicknesses by name: %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.GetSicknessesByNameResponse]{Data: response})
}