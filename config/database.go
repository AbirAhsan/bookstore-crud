package config

import (
	"bookstore-crud/helper"
	"bookstore-crud/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "abir"
	password = "1234"
	dbName   = "test"
)

func ConnectBookDatabase() {
	dns := fmt.Sprintf("host=%s port=%d user= %s  password=%s dbName=%s sslmode=disable", host, port, user, password, dbName)
	database, err := gorm.Open(postgres.Open(dns), &gorm.Config{}) // Use the PostgreSQL driver

	helper.ErrorPanic(err)

	err = database.AutoMigrate(&models.BookModel{})
	if err != nil {
		return
	}

	Db = database
}
