package database

import (
	"localhost/models"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const maxRetries = 5
const retryDelay = 5 * time.Second

var DB *gorm.DB

func buildConnectionString(user, password, host, dbName string) string {
	var connectionString strings.Builder

	connectionString.WriteString(user)
	connectionString.WriteByte(':')
	connectionString.WriteString(password)
	connectionString.WriteString("@tcp(")
	connectionString.WriteString(host)
	connectionString.WriteString(")/")
	connectionString.WriteString(dbName)
	connectionString.WriteString("?charset=utf8&parseTime=True&loc=Local")

	return connectionString.String()
}

func Connect() error {
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	var err error
	var database *gorm.DB
	for i := 0; i < maxRetries; i++ {
		connectionString := buildConnectionString(dbUsername, dbPassword, dbHost, dbName)
		database, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

		if err == nil {
			break
		}

		time.Sleep(retryDelay)
	}

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
