package main

import (
	"log"
	"net/http"
	"os"
	"pdfPro/handler"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if ok != true {
		log.Fatalf("The env variable %s is not set", "PORT")
	}
	http.HandleFunc("/api/v1/genPdf", handler.HandlePdfGen)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
