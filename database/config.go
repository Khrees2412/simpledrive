package database

import (
	"os"

	"simpledrive/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	databaseUrl := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}
	log.Println("Database connected...")

	err = DB.AutoMigrate(models.User{}, models.Folder{}, models.File{})
	if err != nil {
		log.Info("Unable to perform migration")
	}
}
