package api

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/zipf/internal/models"
)

type ZipfBody struct {
	Data string `json:"data"`
}

func HandleZipf(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var b ZipfBody
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	log.Println("decoded body")

	// process data
	l := models.NewLetters(b.Data)
	log.Println("processed data")

	// write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(l)
	log.Println("prepared response")
}

func EnableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}
