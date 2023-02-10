package handler

import (
	"fmt"
	"net/http"
)

func HandlePdfGen(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello from pdf gen")
}
