package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Ammount     uint
	Count       uint
	ExpensesID  uint
}
