package users

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/go-chi/chi/v5"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
    // Extract the user ID from the URL
    userIDStr := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Retrieve the user from the database
    var user models.User
    result := database.DB.First(&user, userID)
    if result.Error != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Respond with the user object
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
