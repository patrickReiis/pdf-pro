package model

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	model "pdfPro/model/entity"
	"time"

	"gorm.io/gorm"
)

const ApiKeyLength int = 80

// Returns a slice with the length of 1
// The element is a timestamp of the time when the this function was called
// The format is ISO 8601
func getRequestTimeStamp() []string {
	return getRequestTimeStampImpl()
}

// Implementation of the `getRequestTimeStamp` function
func getRequestTimeStampImpl() []string {
	return []string{fmt.Sprintf(`%s`, time.Now().UTC().Format(time.RFC3339))}
}

func getRandomApiKey() string {
	return getRandomApiKeyImpl()
}

func getRandomApiKeyImpl() string {
	apiKeyLetters := "abcdefghijklmnopqrstuvwxyz"
	var max = *big.NewInt(int64(len(apiKeyLetters))) // the maximum number allowed when getting a random number
	var apiKey string

	for i := 0; i < ApiKeyLength; i++ {
		randomNumber, _ := rand.Int(rand.Reader, &max) // ignoring potential error
		apiKey += string(apiKeyLetters[int(randomNumber.Int64())])
	}

	return apiKey
}

func doesUserAlreadyExists(email string) (doesUserExists bool) {
	return doesUserExistsImpl(email)
}

func doesUserExistsImpl(email string) (doesUserExists bool) {
	var user model.UserAccount
	result := dbGorm.Where("email = ?", email).First(&user)

	userNotFound := errors.Is(result.Error, gorm.ErrRecordNotFound)
	if userNotFound == true {
		return false
	}

	return true
}