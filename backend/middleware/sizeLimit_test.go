package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouteWithRequestSizeLimit(t *testing.T) {

	makeRequest := func(bodyLength int64) {

		// create byte slice with length of `bodyLength`
		data := make([]byte, bodyLength)

		// create a new reader based on that slice
		body := bytes.NewReader(data)

		// POST request to PDF route
		// The route can essentially be anything since the test is focused on the request Size Limit middleware
		req, err := http.NewRequest("POST", "/api/v1/genPdf", body)

		if err != nil {
			t.Fatalf("Could not create a POST request to the '/api/v1/genPdf' route: %s", err)
		}

		// Record mutations for changes in the response writer
		responseRecorder := httptest.NewRecorder()

		// Set the Request size limit callback function
		handlerRequestSize := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			RouteWithRequestSizeLimit(w, r)
		})

		// Dispatch request to the handler
		handlerRequestSize.ServeHTTP(responseRecorder, req)

		res := responseRecorder.Result()

		// check if the body is greater than the allowed size
		if bodyLength > maxBytesInPostRequest {
			if res.StatusCode != http.StatusRequestEntityTooLarge {
				t.Errorf("got http status code %d, wanted %d", res.StatusCode, http.StatusRequestEntityTooLarge)
			}
		} else {
			if res.StatusCode != http.StatusOK {
				t.Errorf("got http status code %d, wanted %d", res.StatusCode, http.StatusOK)
			}
		}

	}

	var oneMb int64 = 1 * (1 << 20)
	makeRequest(oneMb)

	var fiveMb int64 = 5 * (1 << 20)
	makeRequest(fiveMb)

	var tenMb int64 = 10 * (1 << 20)
	makeRequest(tenMb)

}
