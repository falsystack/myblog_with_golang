package infra

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func SetupDB() *gorm.DB {
	env := os.Getenv("ENV")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var (
		db  *gorm.DB
		err error
	)

	if env == "test" {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
		if err != nil {
			log.Fatalln("Failed to connect to database")
		}
		log.Println("Setup sqlite database")
		return db
	}
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}
	log.Println("Setup postgres database")

	return db
}
