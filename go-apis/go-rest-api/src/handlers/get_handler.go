package handlers

import (
	"encoding/json"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Hello, World! version to"}
	json.NewEncoder(w).Encode(response)
}
