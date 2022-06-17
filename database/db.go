package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	DB = connectDB()
	return DB
}

func connectDB() *gorm.DB {
	var err error
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
