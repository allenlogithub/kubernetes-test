package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	config *viper.Viper
	env    string
)

func InitConfig() {
	config = viper.New()
	env = os.Getenv("EXEC_ENV")

	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")

	if err := config.ReadInConfig(); err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return config
}
