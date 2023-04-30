package authJwt

import (
	"pdfPro/model"
	"testing"
)

func TestVerify(t *testing.T) {
	secretKey, err := model.GetJwtSecret()
	if err != nil {
		t.Fatal("The env variable JWT_SECRET is not set")
		return
	}

	type testVerifyCase struct {
		rawToken  string
		secretKey string
		isValid   bool
	}

	testCases := []testVerifyCase{
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiYXVkIjoic2luZ2xlIiwiZXhwIjo5MDA5MDAwMDB9.BVLaUI8XCeVfyIQq35dOV2hRJWm7GC1a1JFtqhy_Nx8", secretKey, false},
		{"BVLaUI8XCeVfyIQq35dOV2hRJWm7GC1a1JFtqhy_Nx8", secretKey, false},
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0IjoidGVzdGluZyJ9.Nw-7rSzxzfqlvKtH6aa4Ya0ATs4w9TkjL6JTyyaOoY0", "secret-test", true},
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0IjoidGVzdGluZyJ9.Nw-7rSzxzfqlvKtH6aa4Ya0ATs4w9TkjL6JTyyaOoY0", "", false},
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0IjoidGVzdGluZyJ9.iNGNQLw5-xgAfAa0euZMq0DQEiLnlIxEwPNkNyIA6NA", "secret-test", false},
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0IjoidGVzdGluZyJ9.iNGNQLw5-xgAfAa0euZMq0DQEiLnlIxEwPNkNyIA6NA", "secret-test", false},
	}

	for _, e := range testCases {
		_, err := Verify(e.rawToken, &CustomClaims{}, e.secretKey)

		// The token should be valid but an error was returned
		if err != nil && e.isValid == true {
			t.Errorf("Expected %t, but got instead the error: '%s'", e.isValid, err)
		}
	}
}
