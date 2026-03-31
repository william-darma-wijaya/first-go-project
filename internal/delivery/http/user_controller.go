package http

import (
	"first-go-project/internal/model"
	"first-go-project/internal/usecase"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.CreateUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CreateUserResponse]{Data: response})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := &model.UserLoginRequest{}
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Login(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to login user: %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserLoginResponse]{Data: response})
}

func (c *UserController) Logout(ctx *fiber.Ctx) error {
	authHandler := ctx.Get("Authorization")

	if authHandler == "" {
		return fiber.ErrUnauthorized
	}

	token := strings.TrimPrefix(authHandler, "Bearer ")

	resp, err := c.UseCase.Logout(ctx.UserContext(), token)
	if err != nil {
		c.Log.Warnf("Failed to logout user: %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserLogoutResponse]{Data: resp})
}

func (c *UserController) RefreshToken(ctx *fiber.Ctx) error {
	request := &model.RefreshTokenRequest{}
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.RefreshJWTToken(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to refresh JWT token: %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.RefreshTokenResponse]{Data: response})
}

func (c *UserController) GetUserWithAddress(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")


	response, err := c.UseCase.GetUserAndAddress(ctx.UserContext(), idStr)
	if err != nil {
		c.Log.Warnf("Failed to get user and address: %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserWithAddressResponse]{Data: response})
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	request := new(model.UpdateUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.UpdateUser(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to update user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UpdateUserResponse]{Data: response})
}

func (c * UserController) DeleteUser(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	response, err := c.UseCase.DeleteUser(ctx.UserContext(), idStr)
	if err != nil {
		c.Log.Warnf("Failed to delete user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.DeleteUserResponse]{Data: response})
}

func (c *UserController) FindUserByIdWithSicknesses(ctx *fiber.Ctx) error {
	req := new(model.GetUserByIdWithSicknessRequest)
	
	if err := ctx.ParamsParser(req); err != nil {
		c.Log.Warnf("Failed to parse params: %+v", err)
		return fiber.ErrBadRequest
	}

	if err := ctx.QueryParser(req); err != nil {
		c.Log.Warnf("Failed to parse query: %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.FindUserByIdWithSicknesses(ctx.UserContext(), req)
	if err != nil {
		c.Log.Warnf("Failed to retrieve user by ID with sicknesses: %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.GetUserByIdWithSicknessResponse]{Data: response})
}