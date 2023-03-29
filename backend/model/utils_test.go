package model

import "testing"

func TestGetRandomApiKey(t *testing.T) {

	randomAPIs := []string{}
	set := make(map[string]bool)

	// generate an API KEY and append it to the slice
	for i := 0; i < 100000; i++ {
		apiKey := GetRandomApiKey()
		randomAPIs = append(randomAPIs, apiKey)

		if len(apiKey) != ApiKeyLength {
			t.Errorf("The API KEY length should be %d, got %d instead", ApiKeyLength, len(apiKey))
		}
	}

	// iterate over the slice with API KEYS and check if there are duplicated ones
	for _, element := range randomAPIs {
		if set[element] == true {
			t.Errorf("Duplicated API KEY: %s", element)
		}

		set[element] = true
	}

}
