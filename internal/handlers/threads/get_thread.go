package threads

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/go-chi/chi/v5"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
)

func GetThread(w http.ResponseWriter, r *http.Request) {
    // Extract the user ID from the URL
    threadIDStr := chi.URLParam(r, "id")
    threadID, err := strconv.Atoi(threadIDStr)
    if err != nil {
        http.Error(w, "Invalid thread ID", http.StatusBadRequest)
        return
    }

    // Retrieve the user from the database
    var thread models.Thread
    result := database.DB.First(&thread, threadID)
    if result.Error != nil {
        http.Error(w, "Thread not found", http.StatusNotFound)
        return
    }

    // Respond with the user object
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(thread)
}