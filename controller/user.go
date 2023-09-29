package controller

import (
	"errors"
	"strings"

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
	if err = db.DB.Where("email = ?", strings.ToLower(email)).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (uc *UserController) GetUserId(idUser int) (user models.User, err error) {
	if err = db.DB.Where("id = ?", idUser).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (uc *UserController) DelUser(email string) error {
	user, err := uc.GetUser(email)
	if err != nil {
		return err
	}
	if err = db.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (uc *UserController) GetAllUser() (users []models.User, err error) {
	if err = db.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (uc *UserController) PutchUser(data map[string]interface{}) (models.User, error) {
	var user models.User
	email, ok := data["email"].(string)
	idUser := 0
	if !ok {
		idUserGet, ok := data["ID"].(float64)

		if !ok {
			return user, errors.New(configuration.ERROR_UPDATE_USER_EMAIL)
		}
		idUser = int(idUserGet)
	}

	user, err := uc.GetUser(email)
	if err != nil {

		user, err = uc.GetUserId(idUser)
		if err != nil {
			return user, err
		}
	}

	if err := db.DB.Model(&user).Updates(data).Error; err != nil {
		return user, err
	}
	return user, nil

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
