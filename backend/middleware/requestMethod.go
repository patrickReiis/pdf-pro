package middleware

import (
	"fmt"
	"net/http"
)

// Returns false if the request is not of Method POST
func RouteOnlyPostMethod(w http.ResponseWriter, r *http.Request) (ok bool) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		fmt.Fprint(w, `{"error":"Only POST method allowed"}`)
		return false
	}
	return true
}