package comments

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/go-chi/chi/v5"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
)

func GetComment(w http.ResponseWriter, r *http.Request) {
    // Extract the comment ID from the URL
    commentIDStr := chi.URLParam(r, "id")
    commentID, err := strconv.Atoi(commentIDStr)
    if err != nil {
        http.Error(w, "Invalid comment ID", http.StatusBadRequest)
        return
    }

    // Retrieve the user from the database
    var comment models.Comment
    result := database.DB.First(&comment, commentID)
    if result.Error != nil {
        http.Error(w, "Comment not found", http.StatusNotFound)
        return
    }

    // Respond with the user object
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(comment)
}