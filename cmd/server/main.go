package main

import (
	"github.com/craigbwagner/product-catalog-service-go/internal/server"
)

func main() {
	r := server.Router()

	server.Start(r)
}
