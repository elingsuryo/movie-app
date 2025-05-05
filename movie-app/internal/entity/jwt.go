package entity

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	Username string `json:"username"`
	FUllName string `json:"full_name"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}