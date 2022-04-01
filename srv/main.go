package main

import (
	"srv/config"
	"srv/databases"
	"srv/server"
)

func main() {
	config.InitConfig()
	databases.InitMysql()
	server.InitServer()
}
