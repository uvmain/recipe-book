package main

import (
	"encoding/json"
	"log"
	"net/http"
	"recipebook/auth"
	"recipebook/database"
	"recipebook/images"
	"recipebook/logic"
	"recipebook/types"
	"strconv"
	"strings"
	"time"

	"github.com/rs/cors"
)

func enableCdnCaching(w http.ResponseWriter) {
	expiryDate := time.Now().AddDate(1, 0, 0)
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	w.Header().Set("Expires", expiryDate.String())
}

func StartServer() {
	router := http.NewServeMux()

	// frontend
	distDir := http.Dir("../dist")
	fileServer := http.FileServer(distDir)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// serve static files
		if _, err := distDir.Open(r.URL.Path); err == nil {
			enableCdnCaching(w)
			fileServer.ServeHTTP(w, r)
			return
		}
		// serve index.html for non-static files
		http.ServeFile(w, r, "../dist/index.html")
	})

	//auth
	router.HandleFunc("POST /api/login", auth.LoginHandler)
	router.HandleFunc("GET /api/logout", auth.LogoutHandler)
	router.HandleFunc("GET /api/check-session", auth.CheckSessionHandler)

	// standard routes
	router.HandleFunc("GET /api/recipe-count", handleGetRecipeCount)
	router.HandleFunc("GET /api/recipes", handleGetRecipesOrderedByDateCreated)
	router.HandleFunc("GET /api/recipes/{slug}", handleGetRecipeBySlug)
	router.HandleFunc("GET /api/recipecards", handleGetRecipeCardsOrderedByDateCreated)
	router.HandleFunc("GET /api/random-recipe", handleGetRandomRecipe)
	router.HandleFunc("GET /api/images/{filename}", handleGetImageByFilename)

	// authenticated routes
	router.Handle("POST /api/recipes", auth.AuthMiddleware(http.HandlerFunc(handlePostRecipe)))
	router.Handle("PATCH /api/recipes/{slug}", auth.AuthMiddleware(http.HandlerFunc(handlePatchRecipeBySlug)))
	router.Handle("DELETE /api/recipes/{slug}", auth.AuthMiddleware(http.HandlerFunc(handleDeleteRecipeBySlug)))
	router.Handle("POST /api/images", auth.AuthMiddleware(http.HandlerFunc(handlePostNewImage)))
	router.Handle("PATCH /api/images/{filename}", auth.AuthMiddleware(http.HandlerFunc(handlePatchImageByFilename)))
	router.Handle("DELETE /api/images/{filename}", auth.AuthMiddleware(http.HandlerFunc(handleDeleteImageByFilename)))

	handler := cors.AllowAll().Handler(router)

	var serverAddress string
	if logic.IsLocalDevEnv() {
		serverAddress = "localhost:8080"
		log.Println("Application running at https://recipebook.localhost")
	} else {
		serverAddress = ":8080"
		log.Println("Application running at http://localhost:8080")
	}

	http.ListenAndServe(serverAddress, handler)
}

func handleGetRecipeCount(w http.ResponseWriter, r *http.Request) {
	count, _ := database.GetRecipeCount()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(count); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func handleGetRecipesOrderedByDateCreated(w http.ResponseWriter, r *http.Request) {
	recipes, _ := database.GetRecipesOrderedByDateCreated()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recipes); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func handleGetRecipeBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	recipe, _ := database.GetRecipeBySlug(slug)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recipe); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func handleGetRecipeCardsOrderedByDateCreated(w http.ResponseWriter, r *http.Request) {
	searchParam := r.URL.Query().Get("filter")
	caloriesParam := r.URL.Query().Get("calories")
	coursesParam := r.URL.Query().Get("courses")
	countryParam := r.URL.Query().Get("country")
	vegetarianParam := r.URL.Query().Get("vegetarian")

	var filters types.Filters

	filters.Search = strings.Split(searchParam, ",")
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

func handleGetImageByFilename(w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("filename")
	imageBlob, err := images.GetImageByFilename(filename)

	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}
	mimeType := http.DetectContentType(imageBlob)
	enableCdnCaching(w)
	w.Header().Set("Content-Type", mimeType)
	w.WriteHeader(http.StatusOK)
	w.Write(imageBlob)
}

func handlePatchRecipeBySlug(w http.ResponseWriter, r *http.Request) {
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

func handlePostRecipe(w http.ResponseWriter, r *http.Request) {
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

func handleDeleteRecipeBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if err := database.DeleteRecipeBySlug(slug); err != nil {
		http.Error(w, "Failed to delete recipe", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Recipe deleted successfully"))
}

func handleGetRandomRecipe(w http.ResponseWriter, r *http.Request) {
	recipe, _ := database.GetRandomRecipe()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recipe); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func handlePostNewImage(w http.ResponseWriter, r *http.Request) {
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

func handleDeleteImageByFilename(w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("filename")
	if err := images.DeleteImageByFilename(filename); err != nil {
		log.Printf("Failed to delete image: %s", err)
		http.Error(w, "Failed to delete image", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image deleted successfully"))
}

func handlePatchImageByFilename(w http.ResponseWriter, r *http.Request) {
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
