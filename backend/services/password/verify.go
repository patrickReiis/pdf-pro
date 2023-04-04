package password

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Verifies if the plain password is equivalent to the hashed password
// Returns false if they are not equal or if an error happens during the parsing of the hash
func Verify(hashedPass, plainPass string) (isEqual bool) {
	return verifyImpl(hashedPass, plainPass)
}

func verifyImpl(hashedPass, plainPass string) (isEqual bool) {
	parts := strings.Split(hashedPass, "$")

	if len(parts) != 6 {
		return false
	}

	if parts[1] != "argon2id" {
		return false
	}

	if parts[2] != "v=19" {
		return false
	}

	var memory uint32
	var time uint32
	var threads uint8

	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		fmt.Println(err)
		return false
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(parts[4])
	if err != nil {
		fmt.Println(err)
		return false
	}

	hash, err := base64.RawStdEncoding.Strict().DecodeString(parts[5])
	if err != nil {
		fmt.Println(err)
		return false
	}

	keyLength := uint32(len(hash))

	hashNew := argon2.IDKey([]byte(plainPass), salt, time, memory, threads, keyLength)

	// subtle.ConstantTimeCompare() used
	// to help prevent timing attacks
	if subtle.ConstantTimeCompare(hash, hashNew) == 1 {
		return true
	}

	return false
}
