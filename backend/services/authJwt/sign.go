package authJwt

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v4"
)

func Sign(payload json.RawMessage) (string, error) {
	return signImpl(payload)
}

func signImpl(payload json.RawMessage) (string, error) {
	var data jwt.MapClaims

	err := json.Unmarshal(payload, &data)

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("secret key sample, get another password through a new env variable later"))

	return tokenString, err
}
