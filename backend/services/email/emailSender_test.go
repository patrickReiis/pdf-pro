package email

import (
	"os"
	"testing"
)

// Test the `SendMail` function
// Only sends email with a specified body, for now does not attach files
func TestSendEmail(t *testing.T) {

	recipient, ok := os.LookupEnv("RECIPIENT_TEST_EMAIL")
	if ok == false {
		t.Error("Env variables RECIPIENT_TEST_EMAIL is not set")
	}

	to := []string{recipient}
	body := []byte("This is a new body message")
	subject := "Test email sending - PDF PRO"

	err := SendEmail(to, subject, body)
	if err != nil {
		t.Errorf("%s", err)
	}

	err = SendEmail([]string{""}, "", []byte(""))
	if err == nil {
		t.Error("The email sender should have returned an error")
	}
}
