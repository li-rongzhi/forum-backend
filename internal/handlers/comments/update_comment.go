package comments

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/go-chi/chi/v5"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
    "log"
)

func UpdateComment(w http.ResponseWriter, r *http.Request) {
    // Extract the comment ID from the URL
    commentIDStr := chi.URLParam(r, "id")
    commentID, err := strconv.Atoi(commentIDStr)
    if err != nil {
        http.Error(w, "Invalid comment ID", http.StatusBadRequest)
        return
    }

    // Decode the incoming JSON payload into a comment object
    var updatedComment models.Comment
    err = json.NewDecoder(r.Body).Decode(&updatedComment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Fetch the existing Comment from the database
    var existingComment models.Comment
    result := database.DB.First(&existingComment, commentID)
    if result.Error != nil {
        http.Error(w, "Comment not found", http.StatusNotFound)
        return
    }

    // Update the existing Comment with new values
    result = database.DB.Model(&existingComment).Updates(updatedComment)
    if result.Error != nil {
        log.Printf("Failed to update Comment: %v", result.Error)
        http.Error(w, "Failed to update comment", http.StatusInternalServerError)
        return
    }

    // Respond with the updated Comment object
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(existingComment)
}