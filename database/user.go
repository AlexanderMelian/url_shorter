package database

import (
	"errors"
	"localhost/models"

	"gorm.io/gorm"
)

func FindAllUsers() (*[]models.User, error) {
	var users []models.User
	err := DB.Find(&users).Error
	if err != nil {
		return nil, errors.New("could not find users")
	}
	return &users, nil
}

func FindUserById(id uint) (models.User, error) {
	var user models.User
	err := DB.First(&user, id).Error
	if err != nil {
		return models.User{}, errors.New("failed")
	}
	return user, nil
}

func CreateUser(user models.User) error {
	err := DB.Create(&user).Error
	if err != nil {
		return errors.New("failed")
	}
	return nil
}

func FindByUsername(username string) (models.User, error) {
	var user models.User
	result := DB.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}
