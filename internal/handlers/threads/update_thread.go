package threads

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/go-chi/chi/v5"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
    "log"
)

func UpdateThread(w http.ResponseWriter, r *http.Request) {
    // Extract the thread ID from the URL
    threadIDStr := chi.URLParam(r, "id")
    threadID, err := strconv.Atoi(threadIDStr)
    if err != nil {
        http.Error(w, "Invalid thread ID", http.StatusBadRequest)
        return
    }

    // Decode the incoming JSON payload into a thread object
    var updatedThread models.Thread
    err = json.NewDecoder(r.Body).Decode(&updatedThread)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Fetch the existing thread from the database
    var existingThread models.Thread
    result := database.DB.First(&existingThread, threadID)
    if result.Error != nil {
        http.Error(w, "Thread not found", http.StatusNotFound)
        return
    }

    // Update the existing thread with new values
    result = database.DB.Model(&existingThread).Updates(updatedThread)
    if result.Error != nil {
        log.Printf("Failed to update thread: %v", result.Error)
        http.Error(w, "Failed to update thead", http.StatusInternalServerError)
        return
    }

    // Respond with the updated thread object
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(existingThread)
}