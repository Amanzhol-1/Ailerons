package middleware

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
