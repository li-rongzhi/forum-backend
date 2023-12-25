package router

import (
    "sample-go-app/internal/routes"
    "github.com/go-chi/chi/v5"
    "github.com/rs/cors"
)

func Setup() *chi.Mux {
    r := chi.NewRouter()
    // Setup CORS middleware
    corsHandler := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"}, // Adjust as necessary
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        AllowCredentials: true,
        MaxAge:           300,
    })
    r.Use(corsHandler.Handler)

    // Set up routes
    setUpRoutes(r)

    return r
}

func setUpRoutes(r chi.Router) {
    r.Group(routes.GetRoutes()) // Correctly applying the group of routes
}
