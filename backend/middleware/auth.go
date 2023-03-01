package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

// Checks for authorization, if so call the `next` function
func RouteWithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		authHeaderParts := strings.Fields(authHeader)

		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Secret" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"error":"The VALUE of the 'Authorization' header is not in the right format. The correct format should be: Secret <api-key>"}`)

			return
		}
	}
}
