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

	desiredHashMetadata := "$argon2id$v=19$m=65536,t=1,p=4$"

	if hashedPassword[:len(desiredHashMetadata)] != desiredHashMetadata {
		t.Errorf("Expected %s, got %s instead", desiredHashMetadata, hashedPassword[:len(desiredHashMetadata)])
	}
}
