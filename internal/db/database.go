package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s", host, user, dbname, password, sslmode)
	var errOpen error
	DB, errOpen = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errOpen != nil {
		log.Fatalf("ошибка подключения к бд %v", errOpen)
		return errOpen
	}

	//if err := DB.Migrator().DropTable(&Item{}); err != nil {
	//	log.Fatalf("Ошибка дропа: %v", err)
	//}
	//
	//if err := DB.Migrator().DropTable(&Order{}); err != nil {
	//	log.Fatalf("Ошибка дропа: %v", err)
	//}
	//
	//if err := DB.AutoMigrate(&Order{}, &Item{}); err != nil {
	//	log.Fatalf("Ошибка миграции: %v", err)
	//}

	return nil
}
