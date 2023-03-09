package middleware

import (
	"fmt"
	"net/http"
)

// 5mb
const maxBytesInPostRequest = 5 * (1 << 20) // (1 << 20) = 1mb

func RouteWithRequestSizeLimit(w http.ResponseWriter, r *http.Request) (ok bool) {

	if r.ContentLength > int64(maxBytesInPostRequest) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusRequestEntityTooLarge)

		size := maxBytesInPostRequest / (1 << 20)
		fmt.Fprintf(w, `{"error":"The request body exceeded the size of %dmb"}`, size)

		return false
	}
	return true
}
