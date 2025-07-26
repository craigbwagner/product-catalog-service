package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Start(r *chi.Mux) {
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
