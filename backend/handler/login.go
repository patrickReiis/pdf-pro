package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pdfPro/model"
	modelEntity "pdfPro/model/entity"
	"pdfPro/services/authJwt"
	"pdfPro/services/password"
	"time"
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
	// Flow of login:
	// Check email
	// Check password
	// Returns JWT token

	doesUserExists := model.DoesUserAlreadyExists(userAccount.Email)
	if doesUserExists == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		errorMsg := map[string]string{"error": fmt.Sprintf("The user %s does not exist", userAccount.Email)}
		errorMsgJson, _ := json.Marshal(errorMsg) // ignoring potential error
		fmt.Fprint(w, string(errorMsgJson))
		return
	}

	isEqual := password.Verify(model.GetUserPasswordByEmail(userAccount.Email), userAccount.Password)
	if isEqual == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		errorMsg := map[string]string{"error": "The password is incorrect"}
		errorMsgJson, _ := json.Marshal(errorMsg) // ignoring potential error
		fmt.Fprint(w, string(errorMsgJson))
		return
	}

	credentials := make(map[string]interface{})
	credentials["email"] = userAccount.Email
	// Current UTC time + 2 hours
	exp := time.Now().UTC().Add(time.Hour * 2)
	credentials["exp"] = exp.Unix()

	payload, _ := json.Marshal(credentials) // ignoring potential error

	token, err := authJwt.Sign(payload)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		errorMsg := map[string]string{"error": "Could not log you in. Our fault."}
		errorMsgJson, _ := json.Marshal(errorMsg) // ignoring potential error
		fmt.Fprint(w, string(errorMsgJson))
		return
	}

	success := make(map[string]string)
	success["success"] = token
	jsonSuccess, _ := json.Marshal(success) // ignoring potential error

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonSuccess))
}
