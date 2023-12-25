package auth_handler
import (
	"encoding/json"
	"net/http"
	"fmt"
	"gorm.io/gorm"
	"errors"

	"sample-go-app/internal/models"
	"sample-go-app/internal/auth"
	"sample-go-app/internal/database"
)

type LoginResponse struct {
    Token string `json:"token"`
    User  UserData   `json:"user"`
}

type UserData struct {
    UserName string `json:"userName"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginCredentials models.User
	json.NewDecoder(r.Body).Decode(&loginCredentials)
	fmt.Printf("The user request value %v", loginCredentials)

	db := database.DB
	// Find the user by username
    var user models.User
    result := db.Where("user_name = ?", loginCredentials.UserName).First(&user)

	if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprint(w, "Invalid credentials")
        } else {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprint(w, "Error querying the database")
        }
        return
    }
	// Verify the password (assuming passwords are stored hashed)
    if !auth.CheckPasswordHash(loginCredentials.Password, user.Password) {
        w.WriteHeader(http.StatusUnauthorized)
        fmt.Fprint(w, "Invalid credentials")
        return
    }

	// Generate a token if the credentials are valid
    tokenString, err := auth.GenerateToken(user.UserID, user.UserName)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "Error generating the token")
        return
    }

    response := LoginResponse{
        Token:    tokenString,
        User: UserData{
            UserName: user.UserName,
        },
    }

	w.WriteHeader(http.StatusOK)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}