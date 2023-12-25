package categories

import (
    "encoding/json"
    "net/http"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
)

func ListCategories(w http.ResponseWriter, r *http.Request) {
    // Retrieve all categories from the database
    var categories []models.Category
    result := database.DB.Find(&categories)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the list of categories
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(categories)
}