package handler

import (
	"fmt"
	"net/http"
)

func HandlePdfGen(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		fmt.Fprint(w, "Only POST method allowed")
		return
	}

	fmt.Fprint(w, "Hello from pdf gen")
}
