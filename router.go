package main

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"net/http"
	"path"
	"recipebook/core/auth"
	"recipebook/core/config"
	"recipebook/core/handlers"
	"recipebook/core/logic"
	"strings"
	"time"

	"github.com/rs/cors"
)

//go:embed all:frontend/dist
var dist embed.FS
var distSubFS fs.FS
var err error

func StartServer() {
	router := http.NewServeMux()

	distSubFS, err = fs.Sub(dist, "frontend/dist")
	if err != nil {
		log.Fatal("Failed to create sub filesystem:", err)
	}

	router.HandleFunc("/", HandleFrontend)

	//auth
	router.HandleFunc("POST /api/login", auth.LoginHandler)
	router.HandleFunc("GET /api/logout", auth.LogoutHandler)
	router.HandleFunc("GET /api/check-session", auth.CheckSessionHandler)

	// standard routes
	router.HandleFunc("GET /api/recipe-count", handlers.HandleGetRecipeCount)
	router.HandleFunc("GET /api/recipes", handlers.HandleGetRecipesOrderedByDateCreated)
	router.HandleFunc("GET /api/recipes/{slug}", handlers.HandleGetRecipeBySlug)
	router.HandleFunc("GET /api/recipecards", handlers.HandleGetRecipeCardsOrderedByDateCreated)
	router.HandleFunc("GET /api/random-recipe", handlers.HandleGetRandomRecipe)
	router.HandleFunc("GET /api/images/{filename}", handlers.HandleGetImageByFilename)
	router.HandleFunc("GET /api/search", handlers.HandleGetSlugsByFullTextSearch)

	// authenticated routes
	router.Handle("POST /api/recipes", auth.AuthMiddleware(http.HandlerFunc(handlers.HandlePostRecipe)))
	router.Handle("PATCH /api/recipes/{slug}", auth.AuthMiddleware(http.HandlerFunc(handlers.HandlePatchRecipeBySlug)))
	router.Handle("DELETE /api/recipes/{slug}", auth.AuthMiddleware(http.HandlerFunc(handlers.HandleDeleteRecipeBySlug)))
	router.Handle("POST /api/images", auth.AuthMiddleware(http.HandlerFunc(handlers.HandlePostNewImage)))
	router.Handle("PATCH /api/images/{filename}", auth.AuthMiddleware(http.HandlerFunc(handlers.HandlePatchImageByFilename)))
	router.Handle("DELETE /api/images/{filename}", auth.AuthMiddleware(http.HandlerFunc(handlers.HandleDeleteImageByFilename)))

	handler := cors.AllowAll().Handler(router)

	var serverAddress string
	if config.IsLocalDevEnv() {
		serverAddress = "localhost:8080"
		log.Println("Application running at https://recipebook.localhost")
	} else {
		serverAddress = ":8080"
		log.Println("Application running at http://localhost:8080")
	}

	http.ListenAndServe(serverAddress, handler)
}

func HandleFrontend(w http.ResponseWriter, r *http.Request) {
	bootTime := logic.GetBootTime().Truncate(time.Second).UTC()
	lastModified := bootTime.Format(http.TimeFormat)

	cleanPath := path.Clean(r.URL.Path)
	if cleanPath == "/" {
		cleanPath = "/index.html"
	} else {
		cleanPath = strings.TrimPrefix(cleanPath, "/")
	}

	// static content
	file, err := distSubFS.Open(cleanPath)
	if err == nil {
		defer file.Close()

		w.Header().Set("Cache-Control", "public, max-age=0, must-revalidate")
		w.Header().Set("Last-Modified", lastModified)

		ifModifiedSince := r.Header.Get("If-Modified-Since")
		if ifModifiedSince != "" {
			ifTime, err := time.Parse(http.TimeFormat, ifModifiedSince)
			if err == nil && !bootTime.After(ifTime) {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}

		http.ServeContent(w, r, cleanPath, bootTime, file.(io.ReadSeeker))
		return
	}

	// serve index.html for vue-router content
	indexFile, err := distSubFS.Open("index.html")
	if err != nil {
		http.Error(w, "index.html not found", http.StatusNotFound)
		return
	}
	defer indexFile.Close()

	w.Header().Set("Cache-Control", "public, max-age=0, must-revalidate")
	w.Header().Set("Last-Modified", lastModified)

	ifModifiedSince := r.Header.Get("If-Modified-Since")
	if ifModifiedSince != "" {
		ifTime, err := time.Parse(http.TimeFormat, ifModifiedSince)
		if err == nil && !bootTime.After(ifTime) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	http.ServeContent(w, r, "index.html", bootTime, indexFile.(io.ReadSeeker))
}
