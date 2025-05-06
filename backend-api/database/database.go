package database

import (
	"fmt"
	"log"
	"santrikoding/backend-api/config"
	"santrikoding/backend-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// load konfigurasi database dari file .env
	dbUser := config.GetEnv("DB_USER", "root")
	dbPass := config.GetEnv("DB_PASS", "")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "3306")
	dbName := config.GetEnv("DB_NAME", "")

	// format DSN untuk MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	// koneksi ke database
	var error error
	DB, error = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal("Failed to connect to database:", error)
	}

	fmt.Println("Connected to database successfully")

	// auto migrate models
	error = DB.AutoMigrate(&models.User{})
	if error != nil {
		log.Fatal("Failed to migrate database:", error)
	}

	fmt.Println("Database migrated successfully")

}
