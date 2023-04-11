package authJwt

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestSign(t *testing.T) {

	// Current UTC time + 2 hours
	exp := time.Now().UTC().Add(time.Hour * 2)

	data := map[string]interface{}{"hello": "world", "exp": exp.Unix()}

	payload, _ := json.Marshal(data) // ignoring potential error

	payload = json.RawMessage(payload)

	tokenEncoded, err := Sign(payload)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(tokenEncoded)

}
