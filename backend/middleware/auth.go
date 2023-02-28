package middleware

import (
	"net/http"
)

// Checks for authorization, if so call the `next` function
func RouteWithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
		}
	}
}
