package authJwt

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}
