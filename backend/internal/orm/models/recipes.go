package models

import (
	"time"
)

type Ingredient struct{
	ID int
	Name string
	Quantity int
	RecipeId int
	Unit string
	CreatedAt time.Time
}

type Recipe struct{
	ID int
	Title string
	Instructions string
	Ingredients []Ingredient
	ImageUrl string
	Calories int
	Servings int
	PrepTime string
	UserId int
	CategoryId int
	CreatedAt time.Time
}


