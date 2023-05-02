package handler

import (
	"fmt"
	"net/http"
	"pdfPro/model"
	"pdfPro/services/authJwt"
	"strings"
)

func HandleGetApiKey(w http.ResponseWriter, r *http.Request) {
	// There is no validation here since the authen middleware already handles that
	authHeader := r.Header.Get("Authorization")
	authHeaderParts := strings.Fields(authHeader)

	secret, _ := model.GetJwtSecret()

	token, _ := authJwt.Verify(authHeaderParts[1], &authJwt.CustomClaims{}, secret)
	// type assertion
	claims, _ := token.Claims.(*authJwt.CustomClaims)

	apiKey, err := model.GetUserApiKey(claims.Email)
	if err != nil || apiKey == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error":"We couldn't load your API_KEY"}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"apiKey":"%s"}`, apiKey)
}
