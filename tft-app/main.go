package main

import (
	"log"
	"net/http"
	"tft-app/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/challenger", handler.ChallengerHandler)
	mux.HandleFunc("/grandmaster", handler.GrandmasterHandler)
	mux.HandleFunc("/master", handler.MasterHandler)

	// Wrap the mux with CORS middleware
	handler := handler.EnableCORS(mux)

	log.Println("Server is running on :8080")
	http.ListenAndServe(":8080", handler)
}
