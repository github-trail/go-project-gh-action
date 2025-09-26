package routes

import (
	"net/http"

	"go-test-api/src/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/test", handlers.GetHandler).Methods(http.MethodGet)
	return router
}
