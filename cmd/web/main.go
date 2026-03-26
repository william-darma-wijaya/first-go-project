package main

import (
	"first-go-project/internal/config"
	"first-go-project/internal/entity"
	"fmt"
)

func main() {
	viperConfig := config.NewConfigViper()
	viperSecret := config.NewSecretViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	redis := config.NewRedisClient(viperSecret)
	validate := config.NewValidator(viperConfig)
	authConfig := &entity.AuthConfig{
		Secret: viperConfig.GetString("jwt.secret"),
		MinutesExp: viperConfig.GetInt("jwt.minutes_exp"),
	}
	app := config.NewFiber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
		AuthConfig: authConfig,
		RedisClient: redis,
	})

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
