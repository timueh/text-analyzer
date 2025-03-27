package main

import (
	"log"
	"net/http"

	"example.com/zipf/internal/api"
)

func main() {
	log.Println("listening and serving on port 8080")
	http.HandleFunc("/zipf", api.HandleZipf)
	http.ListenAndServe(":8080", nil)
}
