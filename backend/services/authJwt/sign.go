package authJwt

import (
	"encoding/json"
	"pdfPro/model"

	"github.com/golang-jwt/jwt/v4"
)

func Sign(payload json.RawMessage) (string, error) {
	return signImpl(payload)
}

func signImpl(payload json.RawMessage) (string, error) {
	var data jwt.MapClaims

	err := json.Unmarshal(payload, &data)
	if err != nil {
		return "", nil
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	// Get secret key
	key, err := model.GetJwtSecret()
	if err != nil {
		return "", nil
	}

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", nil
	}

	return tokenString, nil
}
