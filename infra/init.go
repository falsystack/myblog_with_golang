package infra

import (
	"github.com/joho/godotenv"
	"log"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error when loading .env file", err)
	}
}
