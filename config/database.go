package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load file env")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create connection to database.")
	}

	//db.AutoMigrate()

	return db
}

// CloseDatabaseConnection method is closing connection database between app and database
func CloseDatabaseConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close connection database.")
	}
	dbSql.Close()
}
