package comments

import (
    "encoding/json"
    "net/http"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
)

func ListComments(w http.ResponseWriter, r *http.Request) {
    // Retrieve all comments from the database
    var comments []models.Comment
    result := database.DB.Find(&comments)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the list of comments
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(comments)
}