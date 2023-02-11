package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandlePdfGen(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Only POST method allowed")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error reading request body")
		return
	}

	var jsonData map[string]interface{}

	err = json.Unmarshal(body, &jsonData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Send valid JSON")
		return
	}

	fmt.Fprintf(w, "Hello from pdf gen %v", jsonData)
}
