package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// NewViper is a function to load config from config.json
// You can change the implementation, for example load from env file, consul, etcd, etc
func NewConfigViper() *viper.Viper {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	// config.AddConfigPath("./../")
	// config.AddConfigPath("./")
	// config.AddConfigPath(".")
	config.AddConfigPath("./cmd/web")
	err := config.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	return config
}


func NewSecretViper() *viper.Viper {
	config := viper.New()

	config.SetConfigName("secret")
	config.SetConfigType("env")
	// config.AddConfigPath("./../")
	// config.AddConfigPath("./")
	// config.AddConfigPath(".")
	config.AddConfigPath("./cmd/web")
	err := config.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error secret file: %w \n", err))
	}

	return config
}