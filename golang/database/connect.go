package database

import (
	"deepsearch/utils"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var db *gorm.DB

func init() {
	config := utils.LoadConfig("./config/server.ini")

	var err error
	db, err = gorm.Open(postgres.Open(config.Database), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&Search{})
}
