package controller

import (
	"errors"

	"github.com/Mau005/MyExpenses/configuration"
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

func (ec *ExpensesController) PuthExpense(data map[string]interface{}) (models.Expenses, error) {
	var expenses models.Expenses
	idExpense, ok := data["ID"].(float64)
	if !ok {
		return expenses, errors.New(configuration.ERROR_EMPTY)
	}

	expense, err := ec.GetExpenses(uint(idExpense))
	if err != nil {
		return expense, err
	}

	//Check Category Expense atribute
	var cc CategoryController
	idCategory, ok := data["IDCategory"].(float64)
	if ok {
		if !(expense.IDCategory == uint(idCategory)) && idCategory != 0 {
			categoryNew, err := cc.GetCategory(uint(idCategory))
			if err != nil {
				return expenses, err
			}
			expense.Category = categoryNew
		}

	}
	//End Check Category

	if err := db.DB.Model(&expense).Updates(data).Error; err != nil {
		return expense, err
	}

	return expense, nil
}

func (ec *ExpensesController) DeleteExpense(idExpenses uint) error {
	expenses, err := ec.GetExpenses(idExpenses)
	if err != nil {
		return err
	}
	if err := db.DB.Delete(&expenses).Error; err != nil {
		return err
	}
	return nil
}

func (ec *ExpensesController) GetExpenses(idExpenses uint) (models.Expenses, error) {
	var expenses models.Expenses
	if err := db.DB.Preload("Product").Preload("Category").Where("id = ?", idExpenses).First(&expenses).Error; err != nil {
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
