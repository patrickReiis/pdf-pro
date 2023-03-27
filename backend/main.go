package main

import (
	"log"
	"net/http"
	"os"
	"pdfPro/handler"
	"pdfPro/middleware"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if ok != true {
		log.Fatalf("The env variable %s is not set", "PORT")
	}

	_, ok = os.LookupEnv("EMAIL")
	if ok != true {
		log.Fatalf("The env variable %s is not set", "EMAIL")
	}

	_, ok = os.LookupEnv("EMAIL_PASSWORD")
	if ok != true {
		log.Fatalf("The env variable %s is not set", "EMAIL_PASSWORD")
	}

	_, ok = os.LookupEnv("EMAIL_HOST")
	if ok != true {
		log.Fatalf("The env variable %s is not set", "EMAIL_HOST")
	}

	_, ok = os.LookupEnv("EMAIL_HOST_PORT")
	if ok != true {
		log.Fatalf("The env variable %s is not set", "EMAIL_HOST_PORT")
	}

	http.HandleFunc("/api/v1/genPdf", middleware.AllMiddleware(handler.HandlePdfGen,
		middleware.MiddlewareRoutes{
			middleware.RouteWithAuth,
			middleware.RouteWithRequestSizeLimit,
			middleware.RouteOnlyPostMethod,
		},
	))

	http.HandleFunc("/api/v1/createAccount", middleware.AllMiddleware(handler.HandleCreateUserAccount,
		middleware.MiddlewareRoutes{
			middleware.RouteWithRequestSizeLimit,
			middleware.RouteOnlyPostMethod,
		},
	))

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
