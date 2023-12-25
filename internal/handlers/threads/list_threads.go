package threads

import (
    "encoding/json"
    "net/http"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
)

func ListThreads(w http.ResponseWriter, r *http.Request) {
    // Retrieve all threads from the database
    var threads []models.Thread
    result := database.DB.Find(&threads)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the list of threads
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(threads)
}