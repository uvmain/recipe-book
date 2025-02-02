package database

import (
	"fmt"
	"log"
	"recipebook/types"
)

func DeleteRecipeBySlug(slug string) error {
	deleteStatement := `DELETE FROM recipes where slug = ?;`
	_, err := Database.Exec(deleteStatement, slug)
	if err != nil {
		log.Printf("Error deleting recipe for %s: %s", slug, err)
		return err
	}
	log.Printf("Metadata row deleted successfully for %s", slug)
	return nil
}

func InsertRecipe(recipe types.Recipe) error {
	stmt, err := Database.Prepare(`INSERT INTO recipes (
		slug, dateCreated, name, author, source, course, country, vegetarian, prepTime, cookingTime, calories, servings, ingredients, instructions, imageFilename, imageWidth, imageHeight
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		recipe.Slug, recipe.DateCreated, recipe.Name, recipe.Author, recipe.Source, recipe.Course,
		recipe.Country, recipe.Vegetarian, recipe.PrepTime, recipe.CookingTime, recipe.Calories,
		recipe.Servings, recipe.Ingredients, recipe.Instructions, recipe.ImageFilename, recipe.ImageWidth, recipe.ImageHeight,
	)
	if err != nil {
		log.Printf("error inserting recipe row: %s", err)
		return err
	}
	log.Printf("Reciperow inserted successfully for %s", recipe.Slug)
	return nil
}

func GetRecipeBySlug(slug string) (types.Recipe, error) {
	var row types.Recipe
	query := `SELECT slug, dateCreated, name, author, source, course, country, vegetarian, prepTime, cookingTime, calories, servings, ingredients, instructions, imageFilename, imageWidth, imageHeight FROM recipes WHERE slug = ?;`

	err := Database.QueryRow(query, slug).Scan(
		&row.Slug, &row.DateCreated, &row.Name, &row.Author, &row.Source, &row.Course, &row.Country, &row.Vegetarian, &row.PrepTime,
		&row.CookingTime, &row.Calories, &row.Servings, &row.Ingredients, &row.Instructions, &row.ImageFilename, &row.ImageWidth, &row.ImageHeight,
	)
	if err != nil {
		return types.Recipe{}, err
	}
	return row, nil
}

func GetRecipesOrderedByDateCreated() ([]types.Recipe, error) {
	var recipes []types.Recipe

	query := `SELECT slug, dateCreated, name, author, source, course, country, vegetarian, prepTime, cookingTime, calories, servings, ingredients, instructions, imageFilename, imageWidth, imageHeight FROM recipes ORDER BY dateCreated DESC;`
	rows, err := Database.Query(query)
	if err != nil {
		log.Printf("Query failed: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe types.Recipe
		if err := rows.Scan(&recipe.Slug, &recipe.DateCreated, &recipe.Name, &recipe.Author, &recipe.Source, &recipe.Course,
			&recipe.Country, &recipe.Vegetarian, &recipe.PrepTime, &recipe.CookingTime, &recipe.Calories, &recipe.Servings,
			&recipe.Ingredients, &recipe.Instructions, &recipe.ImageFilename, &recipe.ImageWidth, &recipe.ImageHeight); err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		return recipes, err
	}

	return recipes, nil
}

func GetRandomRecipe() (types.Recipe, error) {
	var row types.Recipe
	query := `SELECT slug, dateCreated, name, author, source, course, country, vegetarian, prepTime, cookingTime, calories, servings, ingredients, instructions, imageFilename, imageWidth, imageHeight FROM recipes ORDER BY RANDOM() LIMIT 1;`
	err := Database.QueryRow(query).Scan(
		&row.Slug, &row.DateCreated, &row.Name, &row.Author, &row.Source, &row.Course, &row.Country, &row.Vegetarian, &row.PrepTime,
		&row.CookingTime, &row.Calories, &row.Servings, &row.Ingredients, &row.Instructions, &row.ImageFilename, &row.ImageWidth, &row.ImageHeight,
	)
	if err != nil {
		return types.Recipe{}, err
	}
	return row, nil
}

func UpdateRecipeBySlug(slug string, updates map[string]interface{}) error {
	query := "UPDATE recipes SET "
	params := []interface{}{}
	i := 1
	for field, value := range updates {
		query += fmt.Sprintf("%s = ?", field)
		if i < len(updates) {
			query += ", "
		}
		params = append(params, value)
		i++
	}
	query += " WHERE slug = ?"
	params = append(params, slug)

	stmt, err := Database.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(params...)

	if err == nil {
		log.Printf("Recipe updated for %s, %s", slug, updates)
	}
	return err
}

func GetRecipeCardsOrderedByDateCreated() ([]types.RecipeCard, error) {
	var recipeCards []types.RecipeCard

	query := `SELECT slug, dateCreated, name, author, source, course, country, vegetarian, prepTime, cookingTime, calories, servings, imageFilename, imageWidth, imageHeight FROM recipes ORDER BY dateCreated DESC;`
	rows, err := Database.Query(query)
	if err != nil {
		log.Printf("Query failed: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe types.RecipeCard
		if err := rows.Scan(&recipe.Slug, &recipe.DateCreated, &recipe.Name, &recipe.Author, &recipe.Source, &recipe.Course,
			&recipe.Country, &recipe.Vegetarian, &recipe.PrepTime, &recipe.CookingTime, &recipe.Calories, &recipe.Servings,
			&recipe.ImageFilename, &recipe.ImageWidth, &recipe.ImageHeight); err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, err
		}
		recipeCards = append(recipeCards, recipe)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		return recipeCards, err
	}

	return recipeCards, nil
}
