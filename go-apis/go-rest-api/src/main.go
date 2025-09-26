package main

import (
	"log"
	"net/http"

	"go-rest-api/src/routes"
)

func main() {
	router := routes.SetupRoutes()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
