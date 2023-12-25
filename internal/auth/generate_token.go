package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secretKey = []byte("secret-key")

func GenerateToken(user_id uint, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
 	    jwt.MapClaims{
			"user_id": user_id,
			"username": username,
		   	"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
	   return "", err
	}
	return tokenString, nil
}