package storage

import (
	"file-storage/internal/domain/file_info"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var Database DBInstance

func ConnectDb() {
	dsn := "host=localhost user=postgres password=password dbname=file_storage port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the databese. \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connect to the databese successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	db.AutoMigrate(&file_info.FileInfo{})

	Database = DBInstance{Db: db}
}
