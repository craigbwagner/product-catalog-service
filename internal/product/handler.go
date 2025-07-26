package product

import (
	"encoding/json"
	"net/http"
)

func Handler() http.Handler {
	mux :=http.NewServeMux()

	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/products", handleProducts)

	return mux
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		products := []string{"laptop", "phone", "tablet"}
		json.NewEncoder(w).Encode(products)

		return
	}

	http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
}
