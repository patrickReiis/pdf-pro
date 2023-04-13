package authJwt

import (
	"encoding/json"
	"testing"
	"time"
)

func TestSign(t *testing.T) {

	// Current UTC time + 2 hours
	exp := time.Now().UTC().Add(time.Hour * 2)

	data := map[string]interface{}{"hello": "world", "exp": exp.Unix()}

	payload, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	_, err = Sign(payload)
	if err != nil {
		t.Error(err)
	}

}
