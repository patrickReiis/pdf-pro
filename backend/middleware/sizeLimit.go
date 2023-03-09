package middleware

import (
	"fmt"
	"net/http"
)

func RouteWithRequestSizeLimit(w http.ResponseWriter, r *http.Request) (ok bool) {

	// 10mb
	maxBytes := 10 * (1 << 20) // 1 << 20 = 1mb

	if r.ContentLength > int64(maxBytes) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":"The request body exceeded the size of 10mb"}`)
		return false
	}
	return true
}
