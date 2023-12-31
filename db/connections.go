package db

import (
	"errors"
	"fmt"
	"log"

	conf "github.com/Mau005/MyExpenses/configuration"
	"github.com/Mau005/MyExpenses/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func GenerateCategory() {
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

func connectionSqlite(database conf.DataBase, debugMode bool) error {
	var err error

	logDebug := logger.Silent
	if debugMode {
		logDebug = logger.Warn
	}

	DB, err = gorm.Open(sqlite.Open(database.SqlitePath), &gorm.Config{
		Logger: logger.Default.LogMode(logDebug),
	})
	if err != nil {
		return err
	}
	AutoMigrate()
	log.Println(conf.CONNECTION_SQLITE)
	return nil

}

func connectionMysql(database conf.DataBase, debugMode bool) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.User, database.Password, database.Host, database.Port, database.NameDB)

	logDebug := logger.Silent
	if debugMode {
		logDebug = logger.Warn
	}
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logDebug),
	})
	if err != nil {
		return err
	}
	AutoMigrate()
	log.Println(conf.CONNECTION_MYSQL)
	return nil
}

func ConnectionDataBase() error {
	switch conf.Config.DataBase.Engine {
	case "sqlite":
		return connectionSqlite(conf.Config.DataBase, conf.Config.Server.Debug)

	case "mysql":
		return connectionMysql(conf.Config.DataBase, conf.Config.Server.Debug)

	default:
		return errors.New(conf.ERROR_DATABASE_GET)
	}

}

func AutoMigrate() {
	DB.AutoMigrate(&models.Expenses{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.User{})
}
