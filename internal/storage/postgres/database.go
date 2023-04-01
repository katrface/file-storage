package postgres

import (
	"file-storage/internal/domain/file_info"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

func ConnectDB(postgresURL string) {
	db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the databese. \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connect to the databese successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	db.AutoMigrate(&file_info.FileInfo{})

	Database = DBInstance{DB: db}
}

func CloseDB() {
	sqlDB, err := Database.DB.DB()

	if err != nil {
		log.Println(err)
	}

	sqlDB.Close()
}
