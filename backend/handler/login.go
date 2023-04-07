package handler

import (
	"fmt"
	"net/http"
)

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Implement")
}
