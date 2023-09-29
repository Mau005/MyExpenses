package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"-"`
	Names     string `json:"names"`
	LastNames string `json:"last_names"`
	Salary    int    `json:"salary"`
}
