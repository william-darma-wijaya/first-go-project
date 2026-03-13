package route

import (
	"first-go-project/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App               *fiber.App
	UserController    *http.UserController
	AddressController *http.AddressController
}

func (c *RouteConfig) Setup() {
	c.SetupRoute()
}

func (c *RouteConfig) SetupRoute() {
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Get("/api/users/:id", c.UserController.GetUserWithAddress)
	c.App.Patch("/api/users/update", c.UserController.UpdateUser)
	c.App.Delete("/api/users/delete/:id", c.UserController.DeleteUser)

	c.App.Patch("/api/addresses/update", c.AddressController.UpdateAddress)
	c.App.Delete("/api/addresses/delete/:id", c.AddressController.DeleteAddress)
}