package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"recipebook/types"
)

var sessionToken = make(map[string]bool)

func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ADMIN_USER")
	password := os.Getenv("ADMIN_PASSWORD")
	passedUsername := r.FormValue("username")
	passedPassword := r.FormValue("password")
	log.Println("User logging in")
	if passedUsername == username && passedPassword == password {
		token := generateToken()
		sessionToken[token] = true
		http.SetCookie(w, &http.Cookie{
			Name:     "appSession",
			Value:    token,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
			Path:     "/",
			MaxAge:   604800,
		})
		log.Println("Login successful")

		sessionCheck := types.SessionCheck{
			LoggedIn: true,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(sessionCheck); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

	} else {
		log.Println("Login unsuccessful, invalid credentials")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("appSession")
		if err != nil || !sessionToken[cookie.Value] {
			log.Println("Unauthorized access attempt")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("User logging out")
	cookie, err := r.Cookie("appSession")
	if err == nil {
		delete(sessionToken, cookie.Value)
		http.SetCookie(w, &http.Cookie{
			Name:   "appSession",
			Value:  "",
			MaxAge: -1,
			Path:   "/",
		})
	}
	sessionCheck := types.SessionCheck{
		LoggedIn: false,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sessionCheck); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func CheckSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionCheck := types.SessionCheck{
		LoggedIn: false,
	}
	cookie, err := r.Cookie("appSession")
	if err == nil {
		if sessionToken[cookie.Value] {
			sessionCheck.LoggedIn = true
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sessionCheck); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
