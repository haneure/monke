package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// registerIndexRoutes registers routes for the index
func registerIndexRoutes(r *mux.Router) {
	r.HandleFunc("/", indexHandler).Methods("GET")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the index page!")
}
