package controller

import (
	"gorm.io/gorm"
	"testing"
	"toyproject_recruiting_community/entities"
	"toyproject_recruiting_community/infra"
	"toyproject_recruiting_community/repositories"
	"toyproject_recruiting_community/usecases"
)

func setup() *gorm.DB {
	db := infra.SetupDB()
	db.Migrator().DropTable(&entities.User{}, &entities.Post{})
	db.AutoMigrate(&entities.User{}, &entities.Post{})

	setupTestData(db)
	//r := gin.Default()

	//return r
	return db
}

func TestCreateUser(t *testing.T) {
	db := setup()

	repository := repositories.NewAuthRepository(db)
	usecase := usecases.NewAuthUsecase(repository)
	controller := NewAuthController(usecase)

}
