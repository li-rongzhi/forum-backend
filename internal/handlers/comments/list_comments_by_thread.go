package comments

import (
    "fmt"
    "encoding/json"
    "net/http"
    "github.com/go-chi/chi/v5"
    // "gorm.io/gorm"
    "sample-go-app/internal/models"
    "sample-go-app/internal/database"
)


func ListCommentsByThread(w http.ResponseWriter, r *http.Request) {
    // Get the threadID from the URL
    threadID := chi.URLParam(r, "threadID")

    // Convert threadID to the appropriate type (e.g., uint)
    var tid uint
    _, err := fmt.Sscan(threadID, &tid)
    if err != nil {
        http.Error(w, "Invalid thread ID", http.StatusBadRequest)
        return
    }

    // Fetch comments from the database based on threadID
    var comments []models.Comment
    result := database.DB.Where("thread_id = ?", tid).Find(&comments)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Convert the comments to JSON and write the response
    jsonResponse, err := json.Marshal(comments)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}
