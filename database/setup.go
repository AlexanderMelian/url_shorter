package database

import (
	"localhost/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbOptions := os.Getenv("DB_OPTIONS")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?" + dbOptions

	//dsn := "root:pass@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}
	err = database.AutoMigrate(&models.User{}, &models.TinyUrl{})
	if err != nil {
		return err
	}

	DB = database
	return nil
}
