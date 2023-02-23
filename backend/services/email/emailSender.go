// The `email` package provides functionality to send emails
package email

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

// Implementation logic of email sending
func emailSender(to []string, subject string, body []byte) error {

	// the env variables are also checked in the `main.go` file
	email, ok := os.LookupEnv("EMAIL")
	if ok != true {
		log.Fatalf("The env variable %s is not set", "EMAIL")
	}
	password, ok := os.LookupEnv("EMAIL_PASSWORD")
	if ok != true {
		log.Fatalf("The env variable %s is not set", "EMAIL_PASSWORD")
	}
	host, ok := os.LookupEnv("EMAIL_HOST")
	if ok != true {
		log.Fatalf("The env variable %s is not set", "EMAIL_HOST")
	}
	port, ok := os.LookupEnv("EMAIL_HOST_PORT")
	if ok != true {
		log.Fatalf("The env variable %s is not set", "EMAIL_HOST_PORT")
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	auth := smtp.PlainAuth("", email, password, host)

	// the header WITHOUT the body
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n", to[0], subject))

	// full email WITH the body
	fullMsg := append(msg, body...)

	// send email
	err := smtp.SendMail(addr, auth, email, to, fullMsg)

	if err != nil {
		return errors.New(fmt.Sprintf("Could not send the email because of: %s", err))
	}

	return nil

}

// Sends an email
// This is a wrapper function
func SendEmail(to []string, subject string, body []byte) error {
	return emailSender(to, subject, body)
}
