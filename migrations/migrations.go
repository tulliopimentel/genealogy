package migrations

import (
	"example/desafio/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB){
	db.AutoMigrate(models.Person{})
	db.AutoMigrate(models.Family{})
	db.AutoMigrate(models.Relationship{})
}