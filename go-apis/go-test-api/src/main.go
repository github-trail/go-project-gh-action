package main

import (
	"log"
	"net/http"

	"go-test-api/src/routes"
)

func main() {
	router := routes.SetupRoutes()
	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatal(err)
	}
}
