package types

import (
	"database/sql"
	"encoding/json"
	"time"
)

type NullString struct {
	sql.NullString
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return json.Marshal(nil) // Return `null` if invalid
	}
	return json.Marshal(ns.String) // Otherwise, return the string value
}

type Recipe struct {
	Slug          string     `json:"slug"`
	DateCreated   time.Time  `json:"dateCreated"`
	Name          NullString `json:"name"`
	Author        NullString `json:"author"`
	Source        NullString `json:"source"`
	Course        NullString `json:"course"`
	Country       NullString `json:"country"`
	Vegetarian    bool       `json:"vegetarian"`
	PrepTime      NullString `json:"prepTime"`
	CookingTime   NullString `json:"cookingTime"`
	Calories      NullString `json:"calories"`
	Servings      NullString `json:"servings"`
	Ingredients   NullString `json:"ingredients"`
	Instructions  NullString `json:"instructions"`
	ImageFilename NullString `json:"imageFilename"`
	ImageWidth    int        `json:"imageWidth"`
	ImageHeight   int        `json:"imageHeight"`
}

type RecipeCard struct {
	Slug          string     `json:"slug"`
	DateCreated   time.Time  `json:"dateCreated"`
	Name          NullString `json:"name"`
	Author        NullString `json:"author"`
	Source        NullString `json:"source"`
	Course        NullString `json:"course"`
	Country       NullString `json:"country"`
	Vegetarian    bool       `json:"vegetarian"`
	PrepTime      NullString `json:"prepTime"`
	CookingTime   NullString `json:"cookingTime"`
	Calories      NullString `json:"calories"`
	Servings      NullString `json:"servings"`
	ImageFilename NullString `json:"imageFilename"`
	ImageWidth    int        `json:"imageWidth"`
	ImageHeight   int        `json:"imageHeight"`
}

type Filters struct {
	Search     []string
	Calories   int
	Courses    []string
	Country    string
	Vegetarian bool
}
