package authJwt

import "github.com/golang-jwt/jwt/v4"

// Verifies wheter the jwt `token` is valid or not
func Verify(token, secretKey string) (*jwt.Token, error) {
	return verifyImpl(token, secretKey)
}

func verifyImpl(token, secretKey string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
}
