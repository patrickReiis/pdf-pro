package password

import (
	"fmt"
	"math/rand"

	"golang.org/x/crypto/argon2"
)

func Hash(password string) string {
	return hashImpl(password)
}

func hashImpl(password string) string {
	salt := getRandomSalt()
	var time uint32 = 1 // recommended according to https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-argon2-03#section-9.3
	var memory uint32 = 64 * 1024
	var threads uint8 = 4
	var keyLen uint32 = 32

	hashedPassword := string(argon2.IDKey([]byte(password), []byte(salt), time, memory, threads, keyLen))

	return hashedPassword
}

func getRandomSalt() string {
	saltLen := 10
	salt := make([]byte, saltLen)
	_, err := rand.Read(salt)

	if err != nil {
		fmt.Println("error during reading secure random number: ", err)
		return "implement-fallback-for-when-error-in-generating-salt"
	}

	return string(salt)
}
