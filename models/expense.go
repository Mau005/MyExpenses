package models

import (
	"gorm.io/gorm"
)

type Expenses struct {
	gorm.Model
	Name        string
	AmountTotal int
	IdCategory  uint
	Category    Category  `gorm:"foreignKey:IdCategory"`
	Product     []Product `gorm:"foreignKey:ExpensesID"`
}
