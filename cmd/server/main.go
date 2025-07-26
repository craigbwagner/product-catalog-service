package main

import (
	"log"
	"net/http"

	"github.com/craigbwagner/product-catalog-service-go/internal/product"
)

func main() {
	handler := product.Handler()

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
