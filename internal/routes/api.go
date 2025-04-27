package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// registerAPIDataRoutes registers routes for the API data
func registerAPIDataRoutes(r *mux.Router) {
	r.HandleFunc("/api/data", apiDataHandler).Methods("GET")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Some data from the API"
	if _, err := fmt.Fprintln(w, data); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}
