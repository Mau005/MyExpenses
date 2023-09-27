package controller

import (
	"github.com/Mau005/MyExpenses/db"
	"github.com/Mau005/MyExpenses/models"
)

type ExpensesController struct{}

func (ec *ExpensesController) CreateExpenses(expenses models.Expenses) (models.Expenses, error) {
	if err := db.DB.Create(&expenses).Error; err != nil {
		return expenses, err
	}

	return expenses, nil
}

func (ec *ExpensesController) SaveExpenses(expenses models.Expenses) (models.Expenses, error) {
	if err := db.DB.Save(&expenses).Error; err != nil {
		return expenses, err
	}

	return expenses, nil
}

func (ec *ExpensesController) DeleteExpenses(expenses models.Expenses) (models.Expenses, error) {
	if err := db.DB.Delete(&expenses).Error; err != nil {
		return expenses, err
	}
	return expenses, nil
}

func (ec *ExpensesController) GetExpenses(idExpenses uint) (models.Expenses, error) {
	var expenses models.Expenses
	if err := db.DB.Preload("Product").Where("id = ?", idExpenses).Find(&expenses).Error; err != nil {
		return expenses, err
	}
	return expenses, nil
}

func (ec *ExpensesController) GetAllExpenses() (expenses []models.Expenses, err error) {

	if err = db.DB.Preload("Product").Find(&expenses).Error; err != nil {
		return expenses, err
	}
	return expenses, nil
}
