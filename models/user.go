package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string
	Works  string
	Salary string
}
