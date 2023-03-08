// Testing the auth middleware

package middleware

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test the Auth middleware,
// it's successful if the user does not exists or exists
func TestRouteWithAuth(t *testing.T) {

	// POST request to PDF route
	// The route can essentially be anything since the test is focused on the Auth middleware
	req, err := http.NewRequest("POST", "/api/v1/genPdf", nil)

	if err != nil {
		t.Fatalf("Could not create a POST request to the '/api/v1/genPdf' route: %s", err)
	}

	req.Header.Add("Authorization", "Secret password")

	// Record mutations for changes in the response writer
	responseRecorder := httptest.NewRecorder()

	// Representation if the auth middleware went right
	var okRoute bool

	// Set the Auth callback function
	handlerAuth := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		okRoute = RouteWithAuth(w, r)
	})

	// Dispatch request to the handler
	handlerAuth.ServeHTTP(responseRecorder, req)

	res := responseRecorder.Result()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode == http.StatusUnauthorized {
		return
	} else if okRoute == true {
		return
	} else {
		fmt.Println(string(body))
		t.Errorf("got %d, wanted %d || or wanted the `okRoute` to be true\nThe response body is: %s", res.StatusCode, http.StatusUnauthorized, string(body))
	}
}

// Test the Auth middleware is an invalid API key
// An invalid API key is when it contains white spaces
func TestRouteWithAuthInvalidKey(t *testing.T) {

	// POST request to PDF route
	// The route can essentially be anything since the test is focused on the Auth middleware
	req, err := http.NewRequest("POST", "/api/v1/genPdf", nil)

	if err != nil {
		t.Fatalf("Could not create a POST request to the '/api/v1/genPdf' route: %s", err)
	}

	// Invalid API key
	// It's invalid because it can't contain white spaces
	req.Header.Add("Authorization", "Secret password spaces are not allowed")

	// Record mutations for changes in the response writer
	responseRecorder := httptest.NewRecorder()

	// Set the Auth callback function
	handlerAuth := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RouteWithAuth(w, r)
	})

	// Dispatch request to the handler
	handlerAuth.ServeHTTP(responseRecorder, req)

	res := responseRecorder.Result()
	body, _ := io.ReadAll(res.Body)

	if (res.StatusCode == http.StatusBadRequest) == false {
		t.Errorf("got %d, wanted %d", res.StatusCode, http.StatusOK)
	}

	expectedErr := `{"error":"The VALUE of the 'Authorization' header is not in the right format. The correct format should be: Secret <api-key>"}`

	if string(body) != expectedErr {
		t.Errorf("got %s, wanted %s", string(body), expectedErr)
	}

}
