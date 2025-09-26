package routes

import (
	"net/http"

	"go-rest-api/src/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", handlers.GetHandler).Methods(http.MethodGet)
	return router
}
