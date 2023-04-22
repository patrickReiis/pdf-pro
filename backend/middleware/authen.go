package middleware

import (
	"fmt"
	"net/http"
	"pdfPro/model"
	"pdfPro/services/authJwt"
	"strings"
)

// Only allows authenticated users to access the determined endpoint
// If the user is authenticated the `next` function is called
func RouteWithAuthentication(w http.ResponseWriter, r *http.Request) (ok bool) {
	authHeader := r.Header.Get("Authorization")
	authHeaderParts := strings.Fields(authHeader)

	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":"The VALUE of the 'Authorization' header is not in the right format. The correct format should be: Bearer <token>"}`)
		return false
	}

	secret, err := model.GetJwtSecret()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error":"We couldn't log you in. It's out fault."}`)
		return false
	}

	token, err := authJwt.Verify(authHeaderParts[1], secret)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":"You need to have a valid authentication to perform this action."}`)
		return false
	}

	fmt.Println(token.Claims)

	return true
}
