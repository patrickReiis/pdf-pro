package password

import (
	"testing"
)

func TestHash(t *testing.T) {
	plainPassword := "testing321"
	hashedPassword := Hash(plainPassword)

	if hashedPassword == plainPassword {
		t.Error("Error during password hashing. The hashed password is equal to the plain password.")
	}

}
