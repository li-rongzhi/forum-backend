package comments

import (
    "net/http"
    "strconv"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
	"log"
	"github.com/go-chi/chi/v5"
)

// DeleteComment handles DELETE requests to delete a comment
func DeleteComment(w http.ResponseWriter, r *http.Request) {
    // Extract the comment ID from the URL
	commentIDStr := chi.URLParam(r, "id")

    // Convert commentID from string to int
    commentID, err := strconv.Atoi(commentIDStr)
    log.Printf("Attempting to delete comment with ID: %d, original string was '%s'", commentID, commentIDStr)
    if err != nil {
        http.Error(w, "Invalid comment ID: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Connect to the database
    db := database.DB
    if db == nil {
        http.Error(w, "database connection is not established", http.StatusInternalServerError)
        return
    }

    // Delete the comment from the database
    result := db.Delete(&models.Comment{}, commentID)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Check if the comment was actually deleted (result.RowsAffected should be 1)
    if result.RowsAffected == 0 {
        http.Error(w, "Comment not found", http.StatusNotFound)
        return
    }

    // Respond with success
    w.WriteHeader(http.StatusOK) // 200 OK
    w.Write([]byte("Comment deleted successfully"))
}