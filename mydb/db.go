package mydb

import (
	// "fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
	"gorm.io/gorm"
	"echoApi/models"
)

var db *gorm.DB
var err error

func DatabaseInit() {
	db, err = gorm.Open(sqlite.Open("store.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// Migrate the schema
	db.AutoMigrate(&models.Product{}, &models.Cart_Item{}, &models.Category{})
}

func DB() *gorm.DB {
	return db
}
