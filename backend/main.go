package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	database "github.com/justinfranco/cookbook/backend/internal/orm"
	"github.com/justinfranco/cookbook/backend/internal/orm/models"
)

func main() {
	// Setup logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

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
		w.WriteHeader(http.StatusOK)
		slog.Debug("Default route hit", "status", w.Header().Get("status"))
	})
	r.Get("/recipes", func(w http.ResponseWriter, r * http.Request) {
		slog.Debug("Getting recipes")
		recipes := database.GetRecipes()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(recipes)
		slog.Debug("Recipes returned to user")
	})
	r.Post("/recipes", func(w http.ResponseWriter, r * http.Request) {
		// Only allow JSON content type
		ct := r.Header.Get("Content-Type")
		if ct != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Only JSON content type is supported"))
			slog.Debug("User tried to create a recipe with an unsupported content type", "content-type", ct)
			return
		}

		// Set max body size to 1MB
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)

		// Disallow unknown fields
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		// Decode the request body into a recipe struct
		var recipe models.Recipe
		err := decoder.Decode(&recipe)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad request"))
			slog.Error("User tried to create a recipe with a bad request body", "err", err)
			return
		}

		database.CreateRecipe(&recipe)
	})
	http.ListenAndServe(":3000", r)
}