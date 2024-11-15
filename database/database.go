package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance  struct {
	Db *gorm.DB
}

var database DbInstance

func ConnectDb() {
	// Connect to the database
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}

	log.Println("Connected to the database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations...")

	// TODO: Run migrations here

	database = DbInstance{Db: db}
}