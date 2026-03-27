package config

import (
	"first-go-project/internal/CustomValidator"
	"first-go-project/internal/delivery/http"
	"first-go-project/internal/delivery/http/middleware"
	"first-go-project/internal/delivery/http/route"
	"first-go-project/internal/entity"
	"first-go-project/internal/gateway/caching"
	"first-go-project/internal/gateway/messaging"
	"first-go-project/internal/repository"
	"first-go-project/internal/usecase"

	"github.com/IBM/sarama"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB            *gorm.DB
	App           *fiber.App
	Log           *logrus.Logger
	Validate      *validator.Validate
	Config        *viper.Viper
	AuthConfig    *entity.AuthConfig
	RedisClient   *redis.Client
	KafkaProducer sarama.SyncProducer
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	addressRepository := repository.NewAddressRepository(config.Log)
	sicknessRepository := repository.NewSicknessRepository(config.Log)
	userSicknessRepository := repository.NewUserSicknessRepository(config.Log)

	// setup caches
	userCache := caching.NewUserCache(config.RedisClient)

	// setup kafka producers
	var userProducer *messaging.UserProducer
	if config.KafkaProducer != nil {
		userProducer = messaging.NewUserProducer(config.KafkaProducer, config.Log)
	}

	// setup validators
	userValidator := CustomValidator.NewUserValidator(config.Validate)
	addressValidator := CustomValidator.NewAddressValidator(config.Validate)
	sicknessValidator := CustomValidator.NewSicknessValidator(config.Validate)
	userSicknessValidator := CustomValidator.NewUserSicknessValidator(config.Validate)

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, userValidator, config.AuthConfig, userRepository, addressRepository, userCache, userSicknessRepository, userProducer)
	addressUseCase := usecase.NewAddressUseCase(config.DB, config.Log, addressValidator, addressRepository)
	sicknessUseCase := usecase.NewSicknessUsecase(config.DB, config.Log, sicknessValidator, userSicknessRepository, sicknessRepository)
	userSicknessUseCase := usecase.NewUserSicknessUsecase(config.DB, config.Log, userSicknessValidator, userSicknessRepository, sicknessRepository)

	// setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	addressController := http.NewAddressController(addressUseCase, config.Log)
	sicknessController := http.NewSicknessController(sicknessUseCase, config.Log)
	userSicknessController := http.NewUserSicknessController(userSicknessUseCase, config.Log)

	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:                    config.App,
		UserController:         userController,
		AddressController:      addressController,
		SicknessController:     sicknessController,
		UserSicknessController: userSicknessController,
		AuthMiddleware:         authMiddleware,
	}
	routeConfig.Setup()
}
