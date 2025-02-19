package database

import (
	"errors"
	"jwtgen/initializers"
	"jwtgen/models"
)

func CreateUser(user models.User) error {
	err := initializers.DB.Create(&user)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func AuthorizeUser(email string) (models.User, error) {
	var user models.User
	initializers.DB.First(&user, "email =?", email)
	if user.ID == 0 {
		return models.User{}, errors.New("User not found")
	}
	return user, nil
}

func GetUser(email string) (models.User, error) {
	var user models.User
	err := initializers.DB.Where("email=?", email).Find(&user)
	if err.Error != nil {
		return models.User{}, errors.New("user data not found")
	}
	return user, nil
}
