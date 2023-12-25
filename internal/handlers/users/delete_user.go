package users

import (
    "net/http"
    "strconv"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
	"log"
	"github.com/go-chi/chi/v5"
)

// DeleteUser handles DELETE requests to delete a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    // Extract the user ID from the URL
	userIDStr := chi.URLParam(r, "id")

    // Convert userID from string to int
    userID, err := strconv.Atoi(userIDStr)
    log.Printf("Attempting to delete user with ID: %d, original string was '%s'", userID, userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Connect to the database
    db := database.DB
    if db == nil {
        http.Error(w, "database connection is not established", http.StatusInternalServerError)
        return
    }

    // Delete the user from the database
    result := db.Delete(&models.User{}, userID)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Check if the user was actually deleted (result.RowsAffected should be 1)
    if result.RowsAffected == 0 {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Respond with success
    w.WriteHeader(http.StatusOK) // 200 OK
    w.Write([]byte("User deleted successfully"))
}
