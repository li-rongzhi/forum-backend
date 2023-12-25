package users

import (
	"encoding/json"
	"net/http"
    "golang.org/x/crypto/bcrypt"
	"sample-go-app/internal/models"
	"sample-go-app/internal/database"
    "time"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// CreateUser handles POST requests to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
    // Decode the incoming JSON payload into a new User struct
    var newUser models.User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    hashedPassword, err := HashPassword(newUser.Password)
    if err != nil {
        http.Error(w, "Failed to hash the password", http.StatusInternalServerError)
        return
    }
    newUser.Password = hashedPassword

    // Validate the newUser object as needed

    // Connect to the database
    db := database.DB // GetDB no longer returns an error
    if db == nil {
        http.Error(w, "database connection is not established", http.StatusInternalServerError)
        return
    }

    // Create the user in the database
    if result := db.Create(&newUser); result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the created user object, excluding the password
    createdUserResponse := struct {
        UserID    uint      `json:"user_id"`
        UserName  string    `json:"user_name"`
        CreatedAt time.Time `json:"created_at"`
    }{
        UserID:    newUser.UserID,
        UserName:  newUser.UserName,
        CreatedAt: newUser.CreatedAt,
    }

    // Respond with the created user object
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated) // 201 Created
    json.NewEncoder(w).Encode(createdUserResponse)
}

