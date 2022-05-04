package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDb(){
	connection := "root:root@tcp(127.0.0.1:8080)/desafio?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil{
		log.Fatal("error", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxIdleTime(time.Hour)
}

func GetDataBase() *gorm.DB{
	return db
}