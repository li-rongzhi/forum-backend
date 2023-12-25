package users

import (
    "encoding/json"
    "net/http"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
    // Retrieve all users from the database
    var users []models.User
    result := database.DB.Find(&users)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the list of users
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
