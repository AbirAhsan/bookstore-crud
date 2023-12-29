package config

import (
	"bookstore-crud/helper"
	"bookstore-crud/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var BookDb *gorm.DB

const (
	dbPath = "./book.db" // SQLite database file path
)

func ConnectBookDatabase() {
	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{}) // Use the SQLite driver

	helper.ErrorPanic(err)

	err = database.AutoMigrate(&models.BookModel{})
	if err != nil {
		helper.ErrorPanic(err)
	}

	BookDb = database
}
