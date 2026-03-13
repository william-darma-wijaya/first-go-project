package config

import (
	"first-go-project/internal/delivery/http"
	"first-go-project/internal/delivery/http/route"
	"first-go-project/internal/repository"
	"first-go-project/internal/usecase"
	"first-go-project/internal/CustomValidator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	addressRepository := repository.NewAddressRepository(config.Log)

	// setup validators
	userValidator := CustomValidator.NewUserValidator(config.Validate)
	addressValidator := CustomValidator.NewAddressValidator(config.Validate)

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, userValidator, userRepository, addressRepository)
	addressUseCase := usecase.NewAddressUseCase(config.DB, config.Log, addressValidator, addressRepository)

	// setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	addressController := http.NewAddressController(addressUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:               config.App,
		UserController:    userController,
		AddressController: addressController,
	}
	routeConfig.Setup()
}
