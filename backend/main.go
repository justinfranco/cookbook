package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	database "github.com/justinfranco/cookbook/backend/internal/orm"
)

func main() {
	r := chi.NewRouter()

	database.ConnectToDB()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(20 * time.Second))

	 // Set up CORS middleware options
	 // TODO: only have this setup in development environments
    cors := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:8080"}, // Replace with your origin
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300, // Maximum value not ignored by any of major browsers
    })
	r.Use(cors.Handler)


	r.Get("/", func(w http.ResponseWriter, r * http.Request) {
		w.Write([]byte("Hello World!"))
	})
	r.Get("/recipes", func(w http.ResponseWriter, r * http.Request) {
		recipes := database.GetRecipes()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(recipes)
	})
	http.ListenAndServe(":3000", r)
}