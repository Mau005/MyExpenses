package db

import (
	"github.com/Mau005/MyExpenses/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GenerateDefaultAttribute() {
	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 1},
		Name:        "Supermercados",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 2},
		Name:        "Alimentos",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 3},
		Name:        "Transporte",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 4},
		Name:        "Alojamiento",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 5},
		Name:        "Entretenimiento",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 6},
		Name:        "Salud",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 7},
		Name:        "Educación",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 8},
		Name:        "Ropa y accesorios",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 9},
		Name:        "Tecnología",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 10},
		Name:        "Hogar y decoración",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 11},
		Name:        "Viajes",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 12},
		Name:        "Servicios públicos",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 13},
		Name:        "Impuestos",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 14},
		Name:        "Seguros",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 15},
		Name:        "Otros gastos generales",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 16},
		Name:        "Gastos personales",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 17},
		Name:        "Gastos de negocios",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 18},
		Name:        "Regalos y donaciones",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 19},
		Name:        "Ahorros e inversiones",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 20},
		Name:        "Deudas y préstamos",
		Description: "",
	})

	DB.Create(&models.Category{
		Model:       gorm.Model{ID: 21},
		Name:        "Otros",
		Description: "",
	})
}

func ConnectionSqlite() error {
	var err error

	DB, err = gorm.Open(sqlite.Open("MyExpenses.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	AutoMigrate()
	return nil

}

func AutoMigrate() {
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.User{})
}
