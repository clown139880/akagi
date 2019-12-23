package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	var err error
	driver_name := "mysql"
	if driver_name == "" {
		log.Fatal("Invalid driver name")
	}
	dsn := "root:root@tcp(db:3306)/yggdrasil?charset=utf8&parseTime=True&loc=Local"
	if dsn == "" {
		log.Fatal("Invalid DSN")
	}
	DB, err = sqlx.Connect(driver_name, dsn)
	if err != nil {
		log.Fatal(err)
	}
}