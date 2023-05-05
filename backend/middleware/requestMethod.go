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

// Returns false if the request is not of Method GET
func RouteOnlyGetMethod(w http.ResponseWriter, r *http.Request) (ok bool) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		fmt.Fprint(w, `{"error":"Only GET method allowed"}`)
		return false
	}
	return true
}

// Returns false if the request is not of Method DELETE
func RouteOnlyDeleteMethod(w http.ResponseWriter, r *http.Request) (ok bool) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", "DELETE")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		fmt.Fprint(w, `{"error":"Only DELETE method allowed"}`)
		return false
	}
	return true
}
