package authJwt

import "github.com/golang-jwt/jwt/v4"

// Verifies wheter the jwt `token` is valid or not
// `customClaims` is used for parsing the JWT claims (registered claims or user defined claims) in a simple way
// `customClaims` needs to be a pointer otherwise a panic will happen
// See https://pkg.go.dev/github.com/golang-jwt/jwt/v4#ParseWithClaims
func Verify(token string, customClaims jwt.Claims, secretKey string) (*jwt.Token, error) {
	return verifyImpl(token, customClaims, secretKey)
}

func verifyImpl(token string, customClaims jwt.Claims, secretKey string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, customClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
}
