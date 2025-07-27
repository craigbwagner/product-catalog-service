package product

import (
	"encoding/json"
	"net/http"
)

var products = []string{"laptop", "phone", "tablet"}

func List(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}


