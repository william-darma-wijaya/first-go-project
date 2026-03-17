package http

import (
	"first-go-project/internal/model"
	"first-go-project/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserSicknessController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserSicknessUsecase
}

func NewUserSicknessController(useCase *usecase.UserSicknessUsecase, logger *logrus.Logger) *UserSicknessController {
	return &UserSicknessController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *UserSicknessController) CreateNewUserSickness(ctx *fiber.Ctx) error {
	request := new(model.CreateUserSicknessRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse body request: %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.CreateUserSickness(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create new user sickness: %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CreateUserSicknessResponse]{Data: response})
}

func (c *UserSicknessController) CountDiagnosedBySicknessId(ctx *fiber.Ctx) error {
	req := new(model.CountDiagnosedBySicknessIdRequest)
	if err := ctx.ParamsParser(req); err != nil {
		c.Log.Warnf("Failed to parse id params: %+v", err)
		return fiber.ErrBadGateway
	}

	response, err := c.UseCase.CountDiagnosedBySicknessId(ctx.UserContext(), req.Id)
	if err != nil {
		c.Log.Warnf("Failed to count diagnosed by sickness ID: %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CountDiagnosedBySicknessIdResponse]{Data: response})
}
