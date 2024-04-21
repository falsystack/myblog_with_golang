package main

import (
	"toyproject_recruiting_community/entities"
	"toyproject_recruiting_community/infra"
)

func main() {
	infra.Init()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&entities.Post{}, &entities.User{}); err != nil {
		panic("Failed to migrate database")
	}
}
