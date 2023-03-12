package model

import "testing"

// Test if the connection to the database is working
func TestConnectDb(t *testing.T) {
	_, err := connectDb()

	if err != nil {
		t.Error(err)
	}
}
