package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	// Create a new mux router
	r := mux.NewRouter()

	// Register routes for different groups of endpoints
	registerIndexRoutes(r)
	registerAPIDataRoutes(r)

	return r
}

func StartServer(addr string, handler http.Handler) error {
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Printf("Starting server on %s...\n", addr)
	return server.ListenAndServe()
}
