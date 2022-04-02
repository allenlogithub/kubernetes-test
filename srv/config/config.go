package config

import (
	"os"
)

type (
	Config struct {
		Databases Databases
	}

	Databases struct {
		Mysql Mysql
	}

	Mysql struct {
		Address  string
		Net      string
		User     string
		Password string
		DBName   string
	}
)

var (
	config *Config
)

func InitConfig() {
	mysql := Mysql{
		Address:  os.Getenv("MYSQL_ADDRESS"),
		Net:      os.Getenv("MYSQL_NET"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DBNAME"),
	}

	databases := Databases{
		Mysql: mysql,
	}
	config = &Config{
		Databases: databases,
	}
}

func GetConfig() *Config {
	return config
}
