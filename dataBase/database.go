package database

import (
	"log"

	"github.com/Rawan-Temo/Golang-rithimc.git/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger" 
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb (){
    db, err := gorm.Open(sqlite.Open("apiGo.db?_foreign_keys=on"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to Db \n" ,err.Error())
	}

	log.Println("connected to Db Successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")

	db.AutoMigrate(&models.User{} , &models.Product{} ,&models.Order{})

	Database = DbInstance{Db :db}
}