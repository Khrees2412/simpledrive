package database

import (
	"github.com/khrees2412/simpledrive/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func ConnectDB(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Println("Established database connection")

	setDB(db)

	return db, nil
}

func setDB(db *gorm.DB) {
	gormDB = db
}

func DB() *gorm.DB {
	return gormDB
}

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		model.Folder{},
		model.File{},
		model.User{},
	)
}
