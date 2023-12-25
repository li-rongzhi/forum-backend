package threads

import (
    "net/http"
    "strconv"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
	"log"
	"github.com/go-chi/chi/v5"
)

// DeleteThread handles DELETE requests to delete a thread
func DeleteThread(w http.ResponseWriter, r *http.Request) {
    // Extract the thread ID from the URL
	threadIDStr := chi.URLParam(r, "id")

    // Convert threadID from string to int
    threadID, err := strconv.Atoi(threadIDStr)
    log.Printf("Attempting to delete thread with ID: %d, original string was '%s'", threadID, threadIDStr)
    if err != nil {
        http.Error(w, "Invalid thread ID: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Connect to the database
    db := database.DB
    if db == nil {
        http.Error(w, "database connection is not established", http.StatusInternalServerError)
        return
    }

    // Delete the thread from the database
    result := db.Delete(&models.Thread{}, threadID)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Check if the thread was actually deleted (result.RowsAffected should be 1)
    if result.RowsAffected == 0 {
        http.Error(w, "Thread not found", http.StatusNotFound)
        return
    }

    // Respond with success
    w.WriteHeader(http.StatusOK) // 200 OK
    w.Write([]byte("Thread deleted successfully"))
}