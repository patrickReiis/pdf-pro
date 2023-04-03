package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"pdfPro/middleware"
	"pdfPro/model"
	modelEntity "pdfPro/model/entity"
	"strings"
	"testing"

	"gorm.io/gorm"
)

// Test the ability to create an account using the '/api/v1/createAccount' endpoint.
// The test paradigm used in this function is Table Driven Test.
func TestHandleCreateUserAccount(t *testing.T) {
	emailForTesting := os.Getenv("RECIPIENT_TEST_EMAIL")
	if emailForTesting == "" {
		t.Error("RECIPIENT_TEST_EMAIL env variable is not set")
		return
	}

	type testCaseRequest struct {
		requestBody          io.Reader
		expectedStatus       int
		expectedResponseBody string
	}

	var testCases []testCaseRequest

	// This type's purpose is to be converted into a JSON string
	type body map[string]string

	// PASSWORD LENGTH ERROR
	bodyByte, _ := json.Marshal(body{"password": "short pwd", "email": emailForTesting})
	expectedResponse := fmt.Sprintf(`{"error": "Provide a password with a length greater than %d"}`, model.MinimumPasswordLength)
	testCases = append(testCases, testCaseRequest{strings.NewReader(string(bodyByte)), http.StatusBadRequest, expectedResponse})

	// ACCOUNT CREATED SUCCESSFULLY
	bodyByte, _ = json.Marshal(body{"password": "long password.......", "email": emailForTesting})
	expectedResponse = fmt.Sprint(`{"success": "Your account has been created"}`)
	testCases = append(testCases, testCaseRequest{strings.NewReader(string(bodyByte)), http.StatusOK, expectedResponse})

	// INVALID EMAIL
	bodyByte, _ = json.Marshal(body{"password": "long password.......", "not email": "i am not an email"})
	expectedResponse = fmt.Sprint(`{"error": "Provide a valid email"}`)
	testCases = append(testCases, testCaseRequest{strings.NewReader(string(bodyByte)), http.StatusBadRequest, expectedResponse})

	// ERROR READING REQUEST BODY
	expectedResponse = fmt.Sprint(`{"error":"Error reading request body"}`)
	testCases = append(testCases, testCaseRequest{mockReader{}, http.StatusInternalServerError, expectedResponse})

	// INVALID JSON
	expectedResponse = fmt.Sprint(`{"error":"Send valid JSON. The JSON format for creating an account should be {'email': string; 'password': string}"}`)
	testCases = append(testCases, testCaseRequest{strings.NewReader(`i "am ' invalid JSON`), http.StatusBadRequest, expectedResponse})

	// ACCOUNT ALREADY EXISTS
	bodyByte, _ = json.Marshal(body{"password": "long password.......", "email": emailForTesting})
	expectedResponse = fmt.Sprintf(`{"error": "The email '%s' has been taken"}`, emailForTesting)
	testCases = append(testCases, testCaseRequest{strings.NewReader(string(bodyByte)), http.StatusConflict, expectedResponse})

	makeRequest := func(body io.Reader) (respStatus int, respBody string) {
		req, err := http.NewRequest("POST", "/api/v1/createAccount", body)
		if err != nil {
			t.Fatalf("Could not create a POST request to the '/api/v1/createAccount' route: %s", err)
		}

		responseRecorder := httptest.NewRecorder()

		middleware.AllMiddleware(HandleCreateUserAccount,
			middleware.MiddlewareRoutes{
				middleware.RouteWithRequestSizeLimit,
				middleware.RouteOnlyPostMethod,
			},
		).ServeHTTP(responseRecorder, req)

		res := responseRecorder.Result()

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Could not read the response body, error: %s", err)
		}

		return res.StatusCode, string(resBody)
	}

	for i, e := range testCases {
		// Before running the table driven test,
		// the account for testing can't exist (unless the test is of an account that has been taken)
		// to prevent that it's deleted
		_, err := model.DeleteUserAccountByEmail(emailForTesting)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) == false {
				t.Errorf("Could not delete the 'account for testing' before applying the table driven tests, error: %s", err)
				return
			}
		}

		if i == len(testCases)-1 {
			// There is a case where a user tries to register with an email that has been taken
			// To caught this case an account needs to be created first
			// This is the only case where it happens, that's why this case is in the last position of the slice
			var user modelEntity.UserAccount
			user.Email = emailForTesting

			_, err = model.CreateUserAccount(&user)
			if err != nil {
				t.Errorf("Could not create an account to test the 'email has been taken' case, error: %s", err)
				return
			}
		}

		respStatus, respBody := makeRequest(e.requestBody)

		if respStatus != e.expectedStatus {
			t.Errorf("Expected %d, got %d instead", e.expectedStatus, respStatus)
		}

		if respBody != e.expectedResponseBody {
			t.Errorf("Expected %s, got %s instead", e.expectedResponseBody, respBody)
		}
	}
}

// The purpose of this type is for testing
// Returns an error when called the 'Read' method
type mockReader struct{}

func (mr mockReader) Read(p []byte) (int, error) {
	return 0, errors.New("error while reading request body")
}
