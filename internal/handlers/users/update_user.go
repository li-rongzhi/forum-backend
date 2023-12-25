package users

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/go-chi/chi/v5"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
    "log"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    // Extract the user ID from the URL
    userIDStr := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Decode the incoming JSON payload into a user object
    var updatedUser models.User
    err = json.NewDecoder(r.Body).Decode(&updatedUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Fetch the existing user from the database
    var existingUser models.User
    result := database.DB.First(&existingUser, userID)
    if result.Error != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Update the existing user with new values
    result = database.DB.Model(&existingUser).Updates(updatedUser)
    if result.Error != nil {
        log.Printf("Failed to update user: %v", result.Error)
        http.Error(w, "Failed to update user", http.StatusInternalServerError)
        return
    }

    // Respond with the updated user object
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(existingUser)
}
