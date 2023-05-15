package handler

import (
	"fmt"
	"net/http"
	"pdfPro/model"
)

func HandleDeleteUserAccount(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value("userEmail")
	email, ok := userEmail.(string)
	if userEmail == nil || ok == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error":"An error happended, it's out fault."}`)
		return
	}

	_, err := model.DeleteUserAccountByEmail(email)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error":"Couldn't delete your account"}`)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"error":"Your email has been succesfully deleted"}`)
}
