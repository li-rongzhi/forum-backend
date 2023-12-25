package auth

import (
    "github.com/golang-jwt/jwt"
)

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

func GetUserFromToken(tokenString string) (*Claims, error) {
    // Parse the token
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })

    // Check if token is valid
    if err != nil || !token.Valid {
        return nil, err
    }

    return claims, nil
}