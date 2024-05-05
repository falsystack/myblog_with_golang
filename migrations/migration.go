package main

import (
	"toyproject_recruiting_community/entities"
	"toyproject_recruiting_community/infra"
)

func main() {
	infra.Init()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&entities.User{}, &entities.Post{}); err != nil {
		panic("Failed to migrate database")
	}
}
