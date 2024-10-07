package db

import (
	"os"
	"src/index/src/models"
	"src/index/src/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("database_uri")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: false,
		PrepareStmt:                              false,
	})
	if err != nil {
		utils.Logger.Fatalf("Failed to connect to database!")
	}

	// Migrate the schema
	database.AutoMigrate(&models.Hotel{})
	database.AutoMigrate(&models.User{})

	DB = database
}
