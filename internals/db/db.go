package db

import (
	"fmt"
	"kstrategy-backend/internals/models"
	"log"
	"os"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}

	err = DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Println("Failed to migrate")
	}

	log.Println("Connected to MySQL database")
}
