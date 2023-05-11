package model

import (
	"testing"
)

// Test if the connection to the database is working
func TestConnectDb(t *testing.T) {
	err := connectDb()

	if err != nil {
		t.Error(err)
	}
}

// Test if the secret is set
func TestGetJwtSecret(t *testing.T) {
	_, err := GetJwtSecret()
	if err != nil {
		t.Fatalf("%s", err)
	}
}
