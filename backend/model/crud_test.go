package model

import (
	"log"
	"os"
	model "pdfPro/model/entity"
	"testing"
)

func TestCreateUserAccount(t *testing.T) {
	err := InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	emailForTesting := os.Getenv("RECIPIENT_TEST_EMAIL")
	if emailForTesting == "" {
		t.Error("RECIPIENT_TEST_EMAIL env variable is not set")
		return
	}

	doesUserExists := DoesUserAlreadyExists(emailForTesting)
	// The purpose of this test is to test the CREATE operation
	// If the user account already exists delete it so it can be created
	if doesUserExists == true {
		_, err := DeleteUserAccountByEmail(emailForTesting)
		if err != nil {
			t.Errorf("Could not delete the user account so it could be created later, error: %s", err)
			return
		}
	}

	requestsTimestamp := []string{} // empty slice since this test is for creating an account

	user := model.UserAccount{Email: emailForTesting, Password: "create-hash-password-function", RequestsTimestamp: requestsTimestamp, ApiKey: GetRandomApiKey()}

	_, err = CreateUserAccount(&user)
	if err != nil {
		t.Errorf("Could not create an user account, error: %s", err)
	}

}

func TestDeleteUserAccountByEmail(t *testing.T) {
	err := InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	emailForTesting := os.Getenv("RECIPIENT_TEST_EMAIL")
	if emailForTesting == "" {
		t.Error("RECIPIENT_TEST_EMAIL env variable is not set")
		return
	}

	// Main logic for this test:
	// 1 - Check if the test account exists
	// 2 - If does not exist, create it
	// 3 - If it already exists, delete it

	doesUserExists := DoesUserAlreadyExists(emailForTesting)
	if doesUserExists == true {
		_, err := DeleteUserAccountByEmail(emailForTesting)
		if err != nil {
			t.Errorf("Could not delete the account, error: %s", err)
		}
		return
	}

	// Since the user doesn't exist at this point,
	// it neeeds to be created
	// The focus of this test is the DELETING operation, so only the email is needed
	var user model.UserAccount
	user.Email = emailForTesting

	_, err = CreateUserAccount(&user)
	if err != nil {
		t.Errorf("Could not create an account so it could be deleted later, error: %s", err)
		return
	}

	_, err = DeleteUserAccountByEmail(emailForTesting)
	if err != nil {
		t.Errorf("Could not delete the account, error: %s", err)
	}

}
