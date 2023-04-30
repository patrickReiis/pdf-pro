package middleware

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test the authentication middleware,
// it's successful if the user has a valid JWT and the associated user still exists in the database
func TestRouteWithAuthentication(t *testing.T) {

	type testVerifyAuthen struct {
		authHeader     string
		expectedStatus int
		expectedBody   []byte
	}

	testCases := []testVerifyAuthen{
		{
			"Bearer Token Token",
			http.StatusBadRequest,
			[]byte(`{"error":"The VALUE of the 'Authorization' header is not in the right format. The correct format should be: Bearer <token>"}`)},
		{
			"Token",
			http.StatusBadRequest,
			[]byte(`{"error":"The VALUE of the 'Authorization' header is not in the right format. The correct format should be: Bearer <token>"}`)},
		{
			"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0aW5nIjoiaGVsbG8sIHRlc3QifQ.0XqcZDpHNXQCWStAm9VMt9PTNt5kqagZBImk9ddXhps",
			http.StatusBadRequest,
			[]byte(`{"error":"You need to login again"}`)},
	}

	for _, e := range testCases {
		// POST request to delete user
		// The route can essentially be anything since the test is focused on the Authentication middleware
		req, err := http.NewRequest("POST", "/api/v1/deleteUser", nil)

		if err != nil {
			t.Fatalf("Could not create a POST request to the '/api/v1/deleteUser' route: %s", err)
		}

		req.Header.Add("Authorization", e.authHeader)

		// Record mutations for changes in the response writer
		responseRecorder := httptest.NewRecorder()

		// Set the Auth callback function
		handlerAuth := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			RouteWithAuthentication(w, r)
		})

		handlerAuth.ServeHTTP(responseRecorder, req)

		res := responseRecorder.Result()
		body, _ := io.ReadAll(res.Body)

		if res.StatusCode != e.expectedStatus {
			t.Errorf("Got %d, wanted %d\n", res.StatusCode, e.expectedStatus)
		}

		if bytes.Equal(body, e.expectedBody) == false {
			t.Errorf("Got %s, wanted %s\n", string(body), string(e.expectedBody))
		}

	}
}
