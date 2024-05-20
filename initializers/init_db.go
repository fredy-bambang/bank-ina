package initializers

import (
	"log"

	"github.com/fredy-bambang/bank-ina/configs"
	"github.com/fredy-bambang/bank-ina/models"
	"github.com/fredy-bambang/bank-ina/utils"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes database
func InitDB() {
	configs.LoadConfig()
	utils.ConnectSqliteDB()
	DB = utils.DB

	// Migrate the schema
	DB.AutoMigrate(&models.User{}, &models.Task{})
	log.Println("Database init and migrated")

	// return DB
}
