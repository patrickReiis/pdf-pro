package middleware

import (
	"fmt"
	"net/http"
	"pdfPro/model"
	"strings"
)

// Acts as a security middleware so it can guarantee that
// only authorized users can have access to protected routes
// If the user is authorized the `next` function is called
func RouteWithAuth(w http.ResponseWriter, r *http.Request) (ok bool) {
	authHeader := r.Header.Get("Authorization")
	authHeaderParts := strings.Fields(authHeader)

	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Secret" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":"The VALUE of the 'Authorization' header is not in the right format. The correct format should be: Secret <api-key>"}`)
		return false
	}

	apiKey := authHeaderParts[1]

	user := model.GetUserByApiKey(apiKey)

	if user == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"error":"The API KEY you provided is invalid. Check if you are using an API KEY that exists."}`)
		return false
	}
	return true

}
