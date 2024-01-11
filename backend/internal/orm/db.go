package database

import (
	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/justinfranco/cookbook/backend/internal/orm/models"
)

var DB *gorm.DB

func ConnectToDB(){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    os.Getenv("DB_HOST"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_NAME"),
    os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panicMessage := fmt.Sprintf("Unable to connect to the DB, err: %w", err)
		panic(panicMessage)
	}
	fmt.Println("Connected to the DB")
	DB = db
}

func GetRecipes() []models.Recipe {
	var recipes []models.Recipe
	DB.Preload("Ingredients").Find(&recipes)
	return recipes
}

func CreateRecipe(recipe *models.Recipe) {
	slog.Info("Creating recipe")
	tx := DB.Begin()
	tx.Create(&recipe)
	if tx.Error != nil {
		tx.Rollback()
		slog.Error("Error creating recipe", "error", tx.Error)
	}
	tx.Commit()
	slog.Info("Recipe created")
}