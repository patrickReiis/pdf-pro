package authJwt

import (
	"encoding/json"
	"testing"
	"time"
)

func TestSign(t *testing.T) {
	// Tests 2 possible scenarios, the first case should be correct and no errors should happen
	// The second case should return return an error since a string will try to be parsed into a JSON object

	// ------------ FIRST CASE
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

	// ------------ SECOND CASE
	payload, err = json.Marshal("string")
	if err != nil {
		t.Error(err)
	}

	_, err = Sign(payload)
	if err == nil {
		t.Error("Should have returned the error message: cannot unmarshal string into Go value of type jwt.MapClaims")
	}

}
