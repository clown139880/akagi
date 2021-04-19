package models

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//DB 数据库链接
var DB *gorm.DB

func init() {
	//open a db connection
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	dsn := "root:root@tcp(192.168.4.19:3306)/yggdrasil?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "review:123456@tcp(mysql57-review.danatech.cn:3306)/shibinbin_duidaan_duidaan_server_feature_duidaan_1_init_project?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	if os.Getenv("APP_ENV") != "prod" {
		// DB.LogMode(true)
	}

	//Migrate the schema
	//DB.AutoMigrate(&WeightLog{})
}
