package utils

import (
	"fmt"
	"log"

	"github.com/fredy-bambang/bank-ina/configs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectSqliteDB() {
	config := configs.GetConfig()
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc", config.Database.FilePath)

	var err error
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Connected to database")
}
