package database

import (
	"fmt"
	"log"
	"recipebook/logic"
	"recipebook/types"
	"strings"
)

func GetRecipeCount() (int, error) {
	var count int
	query := "select count(*) from recipes;"

	err := Database.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

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

func InsertRecipe(recipe types.RecipeInsert) error {
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

func GetRecipeCardsFilteredAndOrderedByDateCreated(filters types.Filters) ([]types.RecipeCard, error) {
	var recipes []types.Recipe
	var recipeCards []types.RecipeCard

	query := `SELECT slug, dateCreated, name, author, source, course, country, vegetarian, prepTime, cookingTime, calories, servings, ingredients, instructions, imageFilename, imageWidth, imageHeight FROM recipes ORDER BY dateCreated DESC;`
	rows, err := Database.Query(query)
	if err != nil {
		log.Printf("Query failed: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe types.Recipe
		var filterOut = false
		if err := rows.Scan(&recipe.Slug, &recipe.DateCreated, &recipe.Name, &recipe.Author, &recipe.Source, &recipe.Course,
			&recipe.Country, &recipe.Vegetarian, &recipe.PrepTime, &recipe.CookingTime, &recipe.Calories, &recipe.Servings,
			&recipe.Ingredients, &recipe.Instructions, &recipe.ImageFilename, &recipe.ImageWidth, &recipe.ImageHeight); err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, err
		}

		if filters.Calories > 0 {
			recipeCalories, _ := logic.ConvertNullStringToInt(recipe.Calories)
			recipeServings, _ := logic.ConvertNullStringToInt(recipe.Servings)
			if recipeCalories > 0 && recipeServings > 0 {
				if (recipeCalories / recipeServings) > filters.Calories {
					filterOut = true
				}
			} else {
				filterOut = true
			}
		}

		if filters.Country != "" && filters.Country != recipe.Country.String {
			filterOut = true
		}

		if filters.Vegetarian && filters.Vegetarian != recipe.Vegetarian {
			filterOut = true
		}

		if logic.FilteredSliceIsNotEmpty(filters.Courses) {
			filterIn := false
			for _, course := range logic.FilterEmptyStringsFromSlice(filters.Courses) {
				if course == recipe.Course.String {
					filterIn = true
				}
			}
			if !filterIn {
				filterOut = true
			}
		}

		if logic.FilteredSliceIsNotEmpty(filters.Search) {
			if !recipeContainsStrings(recipe, logic.FilterEmptyStringsFromSlice(filters.Search)) {
				filterOut = true
			}
		}

		if !filterOut {
			recipes = append(recipes, recipe)
		}
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		return recipeCards, err
	}

	for _, recipe := range recipes {
		var recipeCard types.RecipeCard

		recipeCard.Author = recipe.Author
		recipeCard.Calories = recipe.Calories
		recipeCard.CookingTime = recipe.CookingTime
		recipeCard.Country = recipe.Country
		recipeCard.Course = recipe.Course
		recipeCard.DateCreated = recipe.DateCreated
		recipeCard.ImageFilename = recipe.ImageFilename
		recipeCard.ImageHeight = recipe.ImageHeight
		recipeCard.ImageWidth = recipe.ImageWidth
		recipeCard.Name = recipe.Name
		recipeCard.PrepTime = recipe.PrepTime
		recipeCard.Servings = recipe.Servings
		recipeCard.Slug = recipe.Slug
		recipeCard.Source = recipe.Source
		recipeCard.Vegetarian = recipe.Vegetarian

		recipeCards = append(recipeCards, recipeCard)
	}

	return recipeCards, nil
}

func recipeContainsStrings(recipe types.Recipe, searchStrings []string) bool {
	recipeStrings := []string{
		recipe.Name.String,
		recipe.Author.String,
		recipe.Source.String,
		recipe.Course.String,
		recipe.Country.String,
		recipe.ImageFilename.String,
	}

	recipeStrings = append(recipeStrings, strings.Split(recipe.Slug, "-")...)
	recipeStrings = append(recipeStrings, strings.Split(recipe.Ingredients.String, " ")...)
	recipeStrings = append(recipeStrings, strings.Split(recipe.Instructions.String, " ")...)

	if recipe.Vegetarian {
		recipeStrings = append(recipeStrings, "vegetarian")
	}

	recipeStrings = logic.StringArraySortUniqueLowercase(recipeStrings)

	for _, searchString := range searchStrings {
		for _, recipeString := range recipeStrings {
			if strings.Contains(recipeString, strings.ToLower(searchString)) {
				return true
			}
		}
	}
	return false
}
