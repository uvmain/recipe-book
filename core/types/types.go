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
	LastModified  time.Time  `json:"lastModified"`
}

type RecipeInsert struct {
	Slug          string `json:"slug"`
	Name          string `json:"name"`
	Author        string `json:"author"`
	Source        string `json:"source"`
	Course        string `json:"course"`
	Country       string `json:"country"`
	Vegetarian    bool   `json:"vegetarian"`
	PrepTime      string `json:"prepTime"`
	CookingTime   string `json:"cookingTime"`
	Calories      string `json:"calories"`
	Servings      string `json:"servings"`
	Ingredients   string `json:"ingredients"`
	Instructions  string `json:"instructions"`
	ImageFilename string `json:"imageFilename"`
	ImageWidth    int    `json:"imageWidth"`
	ImageHeight   int    `json:"imageHeight"`
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
	LastModified  time.Time  `json:"lastModified"`
}

type Filters struct {
	Search     string
	Calories   int
	Courses    []string
	Country    string
	Vegetarian bool
}

type SessionCheck struct {
	LoggedIn bool `json:"loggedIn"`
}
