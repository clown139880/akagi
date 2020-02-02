package models

import (
	"github.com/jinzhu/gorm"
)

//DB 数据库链接
var DB *gorm.DB

func init() {
	//open a db connection
	var err error
	DB, err = gorm.Open("mysql", "root:root@tcp(db:3306)/yggdrasil?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	DB.LogMode(true)
	//Migrate the schema
	//DB.AutoMigrate(&WeightLog{})
}
