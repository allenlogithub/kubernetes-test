package databases

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"

	"srv/config"
)

var (
	connMysql *sql.DB
	err       error
)

func connectMysql(cfg *mysql.Config, dbName string) *sql.DB {
	if dbName != "" {
		cfg.DBName = dbName
	}
	connMysql, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		if dbName != "" {
			log.Fatal("Connect to Mysql.database:" + dbName + "failed.")
		}
		log.Fatal("Connect to Mysql failed.")
	}

	return connMysql
}

func createMysqlDB(conn *sql.DB, dbName string) {
	_, err = conn.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Create Mysql.database:" + dbName + "failed.")
	}
}

func InitMysql() {
	c := config.GetConfig()
	cfg := mysql.Config{
		User:   c.Databases.Mysql.User,
		Passwd: c.Databases.Mysql.Password,
		Net:    c.Databases.Mysql.Net,
		Addr:   c.Databases.Mysql.Address,
	}

	connMysql = connectMysql(&cfg, "")
	defer connMysql.Close()
	createMysqlDB(connMysql, c.Databases.Mysql.DBName)

	connMysql = connectMysql(&cfg, c.Databases.Mysql.DBName)

	q := `
		CREATE TABLE IF NOT EXISTS content (
			id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			content TEXT NOT NULL
		)
	`
	_, err := connMysql.Exec(q)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Create Mysql.table:content failed.")
	}
}

func GetMysql() *sql.DB {
	return connMysql
}
