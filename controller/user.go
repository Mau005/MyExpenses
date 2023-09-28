package controller

import (
	"errors"

	"github.com/Mau005/MyExpenses/configuration"
	"github.com/Mau005/MyExpenses/db"
	"github.com/Mau005/MyExpenses/models"
)

type UserController struct{}

func (uc *UserController) CreateUser(email, password string) (models.User, error) {
	var user models.User
	if email == "" {
		return user, errors.New(configuration.ERROR_EMPTY)

	}
	if password == "" {
		return user, errors.New(configuration.ERROR_EMPTY)
	}

	user.Email = email
	var api ApiController
	user.Password = api.GenerateCryptPassword(password)

	if err := db.DB.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil

}
func (uc *UserController) GetUser(email string) (user models.User, err error) {
	if err = db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (uc *UserController) GetAllUser() (users []models.User, err error) {
	if err = db.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (uc *UserController) LoginUser(email, password string) (user models.User, token string, err error) {
	var ac ApiController

	user, err = uc.GetUser(email)
	if err != nil {
		return user, token, err
	}

	err = ac.CompareCryptPassword(user.Password, password)
	if err != nil {
		return user, token, err
	}
	token, err = ac.GenerateToken(user)

	if err != nil {
		return user, token, err
	}
	return user, token, err

}
