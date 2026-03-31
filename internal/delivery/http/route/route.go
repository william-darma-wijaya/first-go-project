package route

import (
	"first-go-project/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                    *fiber.App
	UserController         *http.UserController
	AddressController      *http.AddressController
	SicknessController     *http.SicknessController
	UserSicknessController *http.UserSicknessController
	AuthMiddleware         fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/login", c.UserController.Login)
	c.App.Post("/api/users/refreshtoken", c.UserController.RefreshToken)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Patch("/api/users/update", c.UserController.UpdateUser)
	c.App.Post("/api/users/logout", c.UserController.Logout)
	c.App.Get("/api/users/userwithsicknesses/:id", c.UserController.FindUserByIdWithSicknesses)
	c.App.Get("/api/users/userwithaddresses:id", c.UserController.GetUserWithAddress)
	c.App.Delete("/api/users/delete/:id", c.UserController.DeleteUser)

	c.App.Patch("/api/addresses/update", c.AddressController.UpdateAddress)
	c.App.Delete("/api/addresses/delete/:id", c.AddressController.DeleteAddress)

	c.App.Post("/api/sickness", c.SicknessController.CreateNewSickness)
	c.App.Get("/api/sicknesses/getnames", c.SicknessController.FindSicknesses)
	c.App.Get("/api/sickness/:id", c.SicknessController.GetSicknessById)
	c.App.Patch("/api/sickness/update", c.SicknessController.UpdateSickness)
	c.App.Delete("/api/sickness/delete/:id", c.SicknessController.DeleteSickness)
	c.App.Get("/api/sicknesswithusers/:id", c.SicknessController.FindSicknessByIdWithUser)

	c.App.Post("/api/usersickness", c.UserSicknessController.CreateNewUserSickness)
	c.App.Get("/api/usersickness/countdiagnosed/:id", c.UserSicknessController.CountDiagnosedBySicknessId)
}
