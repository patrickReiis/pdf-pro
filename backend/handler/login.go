package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	modelEntity "pdfPro/model/entity"
)

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error":"Error reading request body"}`)
		fmt.Println(err)
		return
	}

	var userAccount modelEntity.UserAccount

	err = json.Unmarshal(body, &userAccount)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		errorMsg := map[string]string{"error": `Send valid JSON. The JSON format for logging in should be {'email': string; 'password': string}`}
		errorMsgJson, _ := json.Marshal(errorMsg) // ignoring potential error
		fmt.Fprint(w, string(errorMsgJson))
		return
	}

	fmt.Fprint(w, "Implement")
}
