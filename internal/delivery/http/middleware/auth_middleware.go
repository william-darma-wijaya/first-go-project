package middleware

import (
	"first-go-project/internal/model"
	"first-go-project/internal/usecase"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NewAuth(userUserCase *usecase.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")

		split := strings.Split(authHeader, " ")
		if len(split) != 2 {
			return fiber.ErrUnauthorized
		}

		token := split[1]
		request := &model.VerifyUserRequest{Token: token}
		userUserCase.Log.Debugf("Authorization : %s", request.Token)

		auth, err := userUserCase.Verify(ctx.UserContext(), request)
		if err != nil {
			userUserCase.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		userUserCase.Log.Debugf("User : %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Auth {
	auth, ok := ctx.Locals("auth").(*model.Auth)
	if !ok {
		return nil
	}
	return auth
}
