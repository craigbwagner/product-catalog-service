package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Start(r *chi.Mux) {
	log.Println("Starting server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
