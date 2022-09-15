package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB is the global database connection
var DB *gorm.DB

// ConnectDatabase connects to the database
func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("Could not connect to the database")
	}
	DB = database
	fmt.Println("Database Connected")

	// Migrate the schema
	DB.AutoMigrate(&Book{})
	fmt.Println("Database Migrated")
}
