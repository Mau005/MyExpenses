package models

import (
	"gorm.io/gorm"
)

type Expenses struct {
	gorm.Model
	Name        string
	Description string
	AmountTotal int
	IDCategory  uint
	UserID      uint
	Category    Category  `gorm:"foreignKey:IDCategory"`
	Product     []Product `gorm:"foreignKey:ExpensesID"`
}
