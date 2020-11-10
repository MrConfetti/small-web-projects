package models

import (
	"github.com/jinzhu/gorm"
	// Needed empty import for sqlite3
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB variable to export
var DB *gorm.DB

// ConnectDatabase connects the database
func ConnectDatabase() {
	// try to open sqlite3 connection
	database, err := gorm.Open("sqlite3", "test.db")

	// check if operation succeeded
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migrate model, do this for every model
	database.AutoMigrate(&Book{})

	DB = database
}
