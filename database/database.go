package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"recipes.com/models"
)

var Db *gorm.DB
var dsn = "root:amichel77@tcp(127.0.0.1:3306)/recipes?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDb() {
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	Db.Logger = logger.Default.LogMode(logger.Info)
	Db.AutoMigrate(&models.User{})
}
