package routes

import (
	"net/http"

	"go-rest-api/src/config"
	"go-rest-api/src/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(cfg *config.Config) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", handlers.GetHandler).Methods(http.MethodGet)

	router.HandleFunc("/api/external", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetExternalHandler(w, r, cfg)
	}).Methods(http.MethodGet)

	
	return router
}
