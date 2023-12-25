package routes

import (
    // "net/http"
    "sample-go-app/internal/handlers/users"
    "sample-go-app/internal/handlers/threads"
    "sample-go-app/internal/handlers/comments"
    "sample-go-app/internal/handlers/categories"
    "sample-go-app/internal/handlers/auth_handler"
    "github.com/go-chi/chi/v5"
)

func GetRoutes() func(r chi.Router) {
    return func(r chi.Router) {
        // User routes
        r.Route("/users", UserRoutes)
        // Thread routes
        r.Route("/threads", ThreadRoutes)
        // Comment routes
        r.Route("/comments", CommentRoutes)
        // Category routes
        r.Route("/categories", CategoryRoutes)
        // Auth routes
        r.Route("/auth", AuthRoutes)
    }
}


func UserRoutes(r chi.Router) {
    r.Get("/", users.ListUsers)            // List all users
    r.Get("/{id}", users.GetUser)         // Get a specific user
    r.Post("/", users.CreateUser)         // Create a new user
    r.Delete("/{id}", users.DeleteUser)   // Delete a user
    r.Put("/{id}", users.UpdateUser)      // Update a user
}

func ThreadRoutes(r chi.Router) {
    r.Get("/", threads.ListThreads)         // List all threads
    r.Get("/{id}", threads.GetThread)       // Get a specific thread
    r.Post("/", threads.CreateThread)       // Create a new thread
    r.Delete("/{id}", threads.DeleteThread) // Delete a thread
    r.Put("/{id}", threads.UpdateThread)    // Update a thread

    // Nested route for comments within a specific thread
    r.Route("/{threadID}/comments", func(r chi.Router) {
        r.Get("/", comments.ListCommentsByThread) // List all comments for a specific thread
        // You can add more nested routes here if needed (POST, DELETE, etc.)
    })
}

func CommentRoutes(r chi.Router) {
    r.Get("/", comments.ListComments)         // List all comments
    r.Get("/{id}", comments.GetComment)       // Get a specific comment
    r.Post("/", comments.CreateComment)       // Create a new comment
    r.Delete("/{id}", comments.DeleteComment) // Delete a comment
    r.Put("/{id}", comments.UpdateComment)    // Update a comment
}

func CategoryRoutes(r chi.Router) {
    r.Get("/", categories.ListCategories)
    r.Post("/", categories.CreateCategory)
    r.Delete("/{id}", categories.DeleteCategory)
}

func AuthRoutes(r chi.Router) {
    r.Post("/", auth_handler.LoginHandler)
    r.Get("/", auth_handler.ProtectedHandler)
}