package jwt_services

import (
    "fmt"
    "time"
    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

func GenerateToken(userID string) (string, error) {
    claims := &jwt.StandardClaims{
        ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
        Issuer:    userID,
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) bool {
    claims := &jwt.StandardClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        fmt.Println(err)
        return false
    }

    return token.Valid
}
