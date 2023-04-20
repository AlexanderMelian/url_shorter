package service

import (
	"errors"
	"localhost/models"
	"localhost/setup"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func FindAllUsers() (*[]models.User, bool) {
	db := setup.DB
	var users []models.User
	db.Find(&users)
	return &users, true
}

func FindUserById(id uint) (*models.User, bool) {
	db := setup.DB
	var user models.User
	db.First(&user, id)
	return &user, true
}

func CreateUser(user models.User) bool {
	db := setup.DB
	db.Create(&user)
	return true
}

func FindByUsername(username string) (*models.User, bool) {
	db := setup.DB

	var user models.User
	result := db.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, false
	}
	return &user, true
}

func Login(username string, password string) (*models.User, bool) {
	usu, exist := FindByUsername(username)
	if !exist {
		return nil, false
	}

	err := bcrypt.CompareHashAndPassword([]byte(usu.Password), []byte(password))
	if err != nil {
		return nil, false
	}
	return usu, true
}
