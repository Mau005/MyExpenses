package controller

import (
	"errors"

	"github.com/Mau005/MyExpenses/configuration"
	"github.com/Mau005/MyExpenses/db"
	"github.com/Mau005/MyExpenses/models"
)

type ProductController struct{}

func (cc *ProductController) CreateProduct(product models.Product) (models.Product, error) {
	var exp ExpensesController
	expense, err := exp.GetExpenses(product.ExpensesID) //Identify expense exist
	if err != nil {
		return product, err
	}
	if expense.ID == 0 {
		return product, errors.New(configuration.ERROR_INDEX)
	}

	if product.Name == "" {
		return product, errors.New(configuration.ERROR_EMPTY_FIELD)
	}

	if err := db.DB.Create(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (cc *ProductController) DelProduct(idProduct uint) error {
	product, err := cc.GetProduct(idProduct)

	if err != nil {
		return err
	}

	if err = db.DB.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}

func (cc *ProductController) GetProduct(id uint) (product models.Product, err error) {

	if err = db.DB.Preload("Expenses").Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (cc *ProductController) GetAllProduct() (product []models.Product, err error) {
	if err = db.DB.Find(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (cc *ProductController) PutchProduct(data map[string]interface{}) (models.Product, error) {
	var product models.Product

	idProduct, ok := data["ID"].(float64)

	if !ok {
		return product, errors.New(configuration.ERROR_UPDATE_PRODUCT)
	}
	product, err := cc.GetProduct(uint(idProduct))
	if err != nil {
		return product, errors.New(configuration.ERROR_UPDATE_PRODUCT)
	}

	if err = db.DB.Model(&product).Updates(data).Error; err != nil {
		return product, err
	}

	return product, nil
}
