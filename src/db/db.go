package db

import (
	"fmt"
	"go-chat-service/src/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbInstance *gorm.DB

func Init() error {
	dns := config.GetEnv("DB_URL", "")
	var err error
	DbInstance, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("Failed to connect to database, error: %s", err)
	}
	log.Println("connected")
	DbInstance.Logger = logger.Default.LogMode(logger.Info)

	return nil

}
