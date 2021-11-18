package database

import (
	"fmt"
	"gofiber-gorm-mysql/config"
	"gofiber-gorm-mysql/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

// DB Declare variable for database
var DB *gorm.DB

func ConnectDB() {
	var err error
	portString := config.Config("DB_PORT")
	port, err := strconv.ParseUint(portString, 10, 32)
	if err != nil {
		fmt.Println("Error :", err.Error())
	}
	// Connection URL to connect to mysql database
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), port, config.Config("DB_NAME"))
	// Connect to the DB and init the DB variable
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Connection opened to database")

	// Migrate the database
	err = DB.AutoMigrate(&model.Note{})
	if err != nil {
		return
	}
	fmt.Println("Database Migrated")
}
