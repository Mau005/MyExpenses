package controller

import (
	"errors"

	"github.com/Mau005/MyExpenses/configuration"
	"github.com/Mau005/MyExpenses/db"
	"github.com/Mau005/MyExpenses/models"
)

type CategoryController struct{}

func (cc *CategoryController) CreateCategory(name string) (models.Category, error) {
	category := models.Category{Name: name}

	if category.Name == "" {
		return category, errors.New(configuration.ERROR_EMPTY_FIELD)
	}
	if err := db.DB.Create(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (cc *CategoryController) DelCategory(idCategory uint) error {
	category, err := cc.GetCategory(idCategory)

	if err != nil {
		return err
	}

	if err = db.DB.Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

func (cc *CategoryController) GetCategoryName(nameCategory string) (category models.Category, err error) {

	if err = db.DB.Where("name = ?", nameCategory).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (cc *CategoryController) GetCategory(id uint) (category models.Category, err error) {

	if err = db.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (cc *CategoryController) GetAllCategory() (category []models.Category, err error) {
	if err = db.DB.Find(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (cc *CategoryController) PutchCategory(data map[string]interface{}) (models.Category, error) {
	var category models.Category

	idCategory, ok := data["ID"].(float64)

	if !ok {
		return category, errors.New(configuration.ERROR_UPDATE_CATEGORY)
	}
	category, err := cc.GetCategory(uint(idCategory))
	if err != nil {
		return category, errors.New(configuration.ERROR_UPDATE_CATEGORY)
	}

	if err = db.DB.Model(&category).Updates(data).Error; err != nil {
		return category, err
	}

	return category, nil
}
