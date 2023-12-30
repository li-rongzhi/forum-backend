package threads

import (
    "encoding/json"
    "net/http"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
    "sample-go-app/internal/auth"
    "strconv"
    "strings"
)

// Assuming you have a CategoryID request field as an array of strings
type ThreadRequest struct {
    Title      string   `json:"title"`
    Content    string   `json:"content"`
    CategoryIDs []string `json:"category_ids"`
}

type PublicThread struct {
    ThreadID   uint              `json:"thread_id"`
    Title      string            `json:"title"`
    Content    string            `json:"content"`
    Categories []models.Category `json:"categories"`
}

func CreateThread(w http.ResponseWriter, r *http.Request) {
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
    // Decode the incoming JSON payload
    var req ThreadRequest
    err = json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Prepare the Categories slice
    var categories []models.Category
    for _, catIDStr := range req.CategoryIDs {
        catID, err := strconv.ParseUint(catIDStr, 10, 32)
        if err != nil {
            http.Error(w, "Invalid CategoryID", http.StatusBadRequest)
            return
        }
        categories = append(categories, models.Category{CategoryID: uint(catID)})
    }

    // Create a new Thread object with the UserID from the token
    newThread := models.Thread{
        Title:      req.Title,
        Content:    req.Content,
        UserID:     claims.UserID,
        Categories: categories,
    }

    // Connect to the database and create the thread
    db := database.DB
    if result := db.Create(&newThread); result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the created thread object
    publicThread := PublicThread{
        ThreadID:   newThread.ThreadID,
        Title:      newThread.Title,
        Content:    newThread.Content,
        Categories: newThread.Categories,
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(publicThread)
}
