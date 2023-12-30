package comments

import (
	"encoding/json"
	"net/http"
    "strings"
    "strconv"

	"sample-go-app/internal/models"
	"sample-go-app/internal/database"
    "sample-go-app/internal/auth"
)

type CommentRequest struct {
    Content    string   `json:"content"`
    ThreadID   string   `json:"thread_id"`
}

type PublicComment struct {
    CommentID  uint     `json:"comment_id"`
    Content    string   `json:"content"`
    ThreadID   uint   `json:"thread_id"`
}

// CreateComment handles POST requests to create a new comment
func CreateComment(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    // Extract the token from the Authorization header
    authHeader := r.Header.Get("Authorization")
    tokenString := strings.TrimPrefix(authHeader, "Bearer ")

    // Verify the token and extract the claims
    err := auth.VerifyToken(tokenString)
    if err != nil {
        http.Error(w, "Invalid or missing token", http.StatusUnauthorized)
        return
    }
    // Extract the claims
    claims, err := auth.GetUserFromToken(tokenString)
    if err != nil {
        http.Error(w, "Invalid or missing token", http.StatusUnauthorized)
        return
    }

    // Decode the incoming JSON payload into a new Comment struct
    var req CommentRequest
    err = json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Validate the newComment object as needed

    threadID, err := strconv.ParseUint(req.ThreadID, 10, 32)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        // Handle the error appropriately
    }

    newComment := models.Comment{
        ThreadID:   uint(threadID),
        UserID:     claims.UserID,
        Content:    req.Content,
    }
    // Connect to the database
    db := database.DB // GetDB no longer returns an error
    if db == nil {
        http.Error(w, "database connection is not established", http.StatusInternalServerError)
        return
    }

    // Create the comment in the database
    if result := db.Create(&newComment); result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    publicComment := PublicComment {
        CommentID: newComment.CommentID,
        Content: newComment.Content,
        ThreadID: newComment.ThreadID,
    }

    // Respond with the created comment object
    w.WriteHeader(http.StatusCreated) // 201 Created
    json.NewEncoder(w).Encode(publicComment)
}