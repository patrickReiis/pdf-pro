// The `email` package provides functionality to send emails
package email

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

// Implementation logic of email sending
func emailSender(to []string, subject string, body []byte, pdfToAttach []byte) error {

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

	if len(pdfToAttach) > 0 {
		msg := []byte(fmt.Sprintf("To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: application/pdf\r\n"+
			"Content-Disposition: attachment; filename=\"pdf pro - status.pdf\"\r\n"+
			"Content-Transfer-Encoding: base64\r\n"+
			"\r\n%s\r\n", to, subject, base64.StdEncoding.EncodeToString(pdfToAttach)))

		msg = append(msg, pdfToAttach...)

		// send email
		err := smtp.SendMail(addr, auth, email, to, msg)

		if err != nil {
			return errors.New(fmt.Sprintf("Could not send the email because of: %s", err))
		}
		return nil
	}

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
func SendEmail(to []string, subject string, body []byte, pdfToAttach []byte) error {
	return emailSender(to, subject, body, pdfToAttach)
}
