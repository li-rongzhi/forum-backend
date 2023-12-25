package comments

import (
	"encoding/json"
	"net/http"

	"sample-go-app/internal/models"
	"sample-go-app/internal/database"
)
// CreateComment handles POST requests to create a new comment
func CreateComment(w http.ResponseWriter, r *http.Request) {
    // Decode the incoming JSON payload into a new Comment struct
    var newComment models.Comment
    err := json.NewDecoder(r.Body).Decode(&newComment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Validate the newComment object as needed

    // Connect to the database
    db := database.DB // GetDB no longer returns an error
    if db == nil {
        http.Error(w, "database connection is not established", http.StatusInternalServerError)
        return
    }

    // Create the comment in the database
    if result := db.Create(&newComment); result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the created comment object
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated) // 201 Created
    json.NewEncoder(w).Encode(newComment)
}