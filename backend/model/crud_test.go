package model

import (
	"os"
	model "pdfPro/model/entity"
	"testing"
)

func TestCreateUserAccount(t *testing.T) {

	emailForTesting := os.Getenv("RECIPIENT_TEST_EMAIL")
	if emailForTesting == "" {
		t.Error("RECIPIENT_TEST_EMAIL env variable is not set")
		return
	}

	requestsTimestamp := []string{} // empty slice since this test is for creating an account

	user := model.UserAccount{Email: emailForTesting, Password: "create-hash-password-function", RequestsTimestamp: requestsTimestamp, ApiKey: getRandomApiKey()}

	doesUserExists := doesUserAlreadyExists(emailForTesting)

	if doesUserExists == true {
		return
	}

	_, err := CreateUserAccount(&user)
	if err != nil {
		t.Errorf("Could not create an user account, error: %s", err)
	}

}
