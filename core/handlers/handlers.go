package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"recipebook/core/database"
	"recipebook/core/images"
	"recipebook/core/types"
	"strconv"
	"strings"
	"time"
)

func HandleGetRecipeCount(w http.ResponseWriter, r *http.Request) {
	count, _ := database.GetRecipeCount()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(count); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func HandleGetRecipesOrderedByDateCreated(w http.ResponseWriter, r *http.Request) {
	recipes, _ := database.GetRecipesOrderedByDateCreated()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recipes); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func HandleGetRecipeBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	recipe, _ := database.GetRecipeBySlug(slug)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recipe); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func HandleGetRecipeCardsOrderedByDateCreated(w http.ResponseWriter, r *http.Request) {
	searchParam := r.URL.Query().Get("filter")
	caloriesParam := r.URL.Query().Get("calories")
	coursesParam := r.URL.Query().Get("courses")
	countryParam := r.URL.Query().Get("country")
	vegetarianParam := r.URL.Query().Get("vegetarian")

	var filters types.Filters

	filters.Search = searchParam
	filters.Calories, _ = strconv.Atoi(caloriesParam)
	filters.Courses = strings.Split(coursesParam, ",")
	filters.Country = countryParam
	filters.Vegetarian, _ = strconv.ParseBool(vegetarianParam)

	recipeCards, _ := database.GetRecipeCardsFilteredAndOrderedByDateCreated(filters)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recipeCards); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func HandleGetImageByFilename(w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("filename")
	imageBlob, lastModified, err := images.GetImageByFilename(filename)
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	ifModifiedSince := r.Header.Get("If-Modified-Since")
	if ifModifiedSince != "" {
		ifTime, err := time.Parse(http.TimeFormat, ifModifiedSince)
		if err == nil && !lastModified.Truncate(time.Second).After(ifTime) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	mimeType := http.DetectContentType(imageBlob)
	w.Header().Set("Content-Type", mimeType)
	w.Header().Set("Last-Modified", lastModified.Truncate(time.Second).UTC().Format(http.TimeFormat))
	w.Header().Set("Cache-Control", "public, max-age=0, must-revalidate")
	w.WriteHeader(http.StatusOK)
	w.Write(imageBlob)
}

func HandlePatchRecipeBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := database.UpdateRecipeBySlug(slug, updates); err != nil {
		log.Printf("Failed to update recipe: %s", err)
		http.Error(w, "Failed to update recipe", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Recipe updated successfully"))
}

func HandlePostRecipe(w http.ResponseWriter, r *http.Request) {
	var updates types.RecipeInsert
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		log.Printf("Invalid JSON: %s", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := database.InsertRecipe(updates); err != nil {
		http.Error(w, "Failed to post recipe", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Recipe posted successfully"))
}

func HandleDeleteRecipeBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if err := database.DeleteRecipeBySlug(slug); err != nil {
		http.Error(w, "Failed to delete recipe", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Recipe deleted successfully"))
}

func HandleGetRandomRecipe(w http.ResponseWriter, r *http.Request) {
	recipe, _ := database.GetRandomRecipe()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recipe); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func HandlePostNewImage(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file: "+err.Error(), http.StatusBadRequest)
		return
	}
	filename := r.FormValue("filename")
	defer file.Close()
	err = images.UploadImage(file, filename)
	if err != nil {
		http.Error(w, "Failed to upload image", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image uploaded successfully"))
}

func HandleDeleteImageByFilename(w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("filename")
	if err := images.DeleteImageByFilename(filename); err != nil {
		log.Printf("Failed to delete image: %s", err)
		http.Error(w, "Failed to delete image", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image deleted successfully"))
}

func HandlePatchImageByFilename(w http.ResponseWriter, r *http.Request) {
	// Delete the existing image
	if err := images.DeleteImageByFilename(r.PathValue("filename")); err != nil {
		log.Printf("Failed to delete image: %s", err)
	}
	// Create the new image
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	filename := r.FormValue("filename")
	if err := images.UploadImage(file, filename); err != nil {
		http.Error(w, "Failed to upload image", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image replaced successfully"))
}

func HandleGetSlugsByFullTextSearch(w http.ResponseWriter, r *http.Request) {
	searchText := r.URL.Query().Get("search-text")
	slugs, _ := database.FullTextSearch(searchText)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(slugs); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
