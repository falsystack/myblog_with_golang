package infra

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Init() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatalln("Error when loading .env file", err)
	}
	log.Println(os.Getenv("ENV"), " IS LOADED")
}
