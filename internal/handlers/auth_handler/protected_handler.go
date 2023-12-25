package auth_handler
import (
	"net/http"
	"fmt"
	"sample-go-app/internal/auth"
	"strings"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
	  w.WriteHeader(http.StatusUnauthorized)
	  fmt.Fprint(w, "Missing authorization header")
	  return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	err := auth.VerifyToken(tokenString)
	if err != nil {
	  w.WriteHeader(http.StatusUnauthorized)
	  fmt.Fprint(w, "Invalid token")
	  return
	}
	claims, err := auth.GetUserFromToken(tokenString)
	if err != nil {
        // Handle error - the token is invalid or not provided
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
	fmt.Printf("%d", claims.UserID)
	fmt.Fprint(w, "Welcome to the the protected area")
}