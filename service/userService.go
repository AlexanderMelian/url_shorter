package service

import (
	"errors"
	"localhost/database"
	"localhost/models"

	"golang.org/x/crypto/bcrypt"
)

func FindAllUsers() (*[]models.User, error) {
	users, err := database.FindAllUsers()
	if err != nil {
		return nil, errors.New("could not find users")
	}

	return users, nil
}

func FindUserById(id uint) (models.User, error) {
	user, err := database.FindUserById(id)
	if err != nil {
		return models.User{}, errors.New("failed")
	}

	return user, nil
}

func CreateUser(user models.User) error {
	err := database.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func FindByUsername(username string) (models.User, error) {
	user, err := database.FindByUsername(username)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func Login(username string, password string) (models.User, error) {
	user, err := database.FindByUsername(username)

	if err != nil {
		return models.User{}, errors.New("failed find by username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("failed bcrypt")
	}
	return user, nil
}
