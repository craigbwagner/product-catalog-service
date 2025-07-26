package product

import (
	"encoding/json"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	products := []string{"laptop", "phone", "tablet"}
	json.NewEncoder(w).Encode(products)
}
