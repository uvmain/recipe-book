package types

import (
	"time"
)

type Recipe struct {
	Slug          string    `json:"slug"`
	DateCreated   time.Time `json:"dateCreated"`
	Name          string    `json:"name"`
	Author        string    `json:"author"`
	Source        string    `json:"source"`
	Course        string    `json:"course"`
	Country       string    `json:"country"`
	Vegetarian    bool      `json:"vegetarian"`
	PrepTime      string    `json:"prepTime"`
	CookingTime   string    `json:"cookingTime"`
	Calories      string    `json:"calories"`
	Servings      string    `json:"servings"`
	Ingredients   string    `json:"Ingredients"`
	Instructions  string    `json:"instructions"`
	ImageFilename string    `json:"imageFilename"`
	ImageWidth    int       `json:"ImageWidth"`
	ImageHeight   int       `json:"imageHeight"`
}

type RecipeCard struct {
	Slug          string    `json:"slug"`
	DateCreated   time.Time `json:"dateCreated"`
	Name          string    `json:"name"`
	Author        string    `json:"author"`
	Source        string    `json:"source"`
	Course        string    `json:"course"`
	Country       string    `json:"country"`
	Vegetarian    bool      `json:"vegetarian"`
	PrepTime      string    `json:"prepTime"`
	CookingTime   string    `json:"cookingTime"`
	Calories      string    `json:"calories"`
	Servings      string    `json:"servings"`
	ImageFilename string    `json:"imageFilename"`
	ImageWidth    int       `json:"ImageWidth"`
	ImageHeight   int       `json:"imageHeight"`
}
