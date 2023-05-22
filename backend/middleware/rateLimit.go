package middleware

import (
	"fmt"
	"net/http"
	"pdfPro/model"
	"strings"
)

func RouteWithRateLimiting(w http.ResponseWriter, r *http.Request) (ok bool) {
	authHeader := r.Header.Get("Authorization")
	authHeaderParts := strings.Fields(authHeader)

	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Secret" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":"The VALUE of the 'Authorization' header is not in the right format. The correct format should be: Secret <api-key>"}`)
		return
	}

	apiKey := authHeaderParts[1]

	timestamps, err := model.GetUserTimestampByApiKey(apiKey)
	if err != nil {
		fmt.Fprint(w, "not found")
		return false
	}

	fmt.Println(timestamps)
	// If you want to implement this middleware:

	model.GetRequestTimeStamp() // call this function to get a list of dates
	// Compare if the request is allowed. For example if it's 1000 requests per hour you can check that
	model.UpdateUserTimestamp() // Go inside this function to implement a mechanism to increment a new timestamp

	return false // return true
}
