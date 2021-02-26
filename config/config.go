package config

import (
	"fmt"
	"project/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB shared gorm.DB object accross the code to use
var DB *gorm.DB

// InitDB to initialize database connection
func InitDB() {

	DB_Username := "root"
	DB_Password := "mysql"
	DB_Port := "3306"
	DB_Host := "127.0.0.1"
	DB_Name := "altastore"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_Username,
		DB_Password,
		DB_Host,
		DB_Port,
		DB_Name)

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

// InitMigrate to initialize database migration
func InitMigrate() {
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Categories{})
	DB.AutoMigrate(&models.Products{})
	DB.AutoMigrate(&models.Transactions{})
}
