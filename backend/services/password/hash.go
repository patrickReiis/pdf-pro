package password

import (
	"encoding/base64"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/argon2"
)

func Hash(password string) string {
	return hashImpl(password)
}

func hashImpl(password string) string {
	salt, _ := getRandomSalt()
	var time uint32 = 1 // recommended according to https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-argon2-03#section-9.3
	var memory uint32 = 64 * 1024
	var threads uint8 = 4
	var keyLen uint32 = 32
	saltB64 := base64.RawStdEncoding.EncodeToString(salt)

	password = base64.RawStdEncoding.EncodeToString(argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen))

	// standard hash format
	hashedPassword := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, time, threads, saltB64, password)

	return hashedPassword
}

func getRandomSalt() ([]byte, error) {
	saltLen := 10
	salt := make([]byte, saltLen)
	_, err := rand.Read(salt)

	if err != nil {
		fmt.Println("error during reading secure random number: ", err)
		return nil, err
	}

	return salt, nil
}
