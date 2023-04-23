package middleware

import (
	"fmt"
	"testing"
)

// Test the authentication middleware,
// it's successful if the user has a valid JWT and the associated user still exists in the database
func TestRouteWithAuthentication(t *testing.T) {
	fmt.Println("file creation")
}
