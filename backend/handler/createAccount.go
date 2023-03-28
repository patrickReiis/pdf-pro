package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pdfPro/model"
	modelEntity "pdfPro/model/entity"
	"pdfPro/services/password"
)

func HandleCreateUserAccount(w http.ResponseWriter, r *http.Request) {

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

		errorMsg := map[string]string{"error": `Send valid JSON. The JSON format for creating an account should be {'email': string; 'password': string}`}
		errorMsgJson, _ := json.Marshal(errorMsg) // ignoring potential error
		fmt.Fprint(w, string(errorMsgJson))
		return
	}

	if userAccount.Email == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error": "Provide a valid email"}`)
		return
	}

	if userAccount.Password == "" || len(userAccount.Password) < model.MinimumPasswordLength {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Provide a password with a length greater than %d"}`, model.MinimumPasswordLength)
		return
	}

	doesUserExists := model.DoesUserAlreadyExists(userAccount.Email)
	if doesUserExists == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "The email '%s' has been taken"}`, userAccount.Email)
		return
	}

	userAccount.Password = password.Hash(userAccount.Password)
	userAccount.ApiKey = model.GetRandomApiKey()

	_, err = model.CreateUserAccount(&userAccount)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Couldn't create your account"}`)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"success": "Your account has been created"}`)

}
