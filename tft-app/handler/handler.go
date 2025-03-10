package handler

import (
	"encoding/json"
	"net/http"
	"tft-app/service"
)

// Helper function to send JSON response
func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// ChallengerHandler - Handles the /tft/league/v1/challenger endpoint
func ChallengerHandler(w http.ResponseWriter, r *http.Request) {
	// Call the service layer to get data from Riot API
	data, err := service.FetchAllChallengers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response
	sendJSONResponse(w, http.StatusOK, data)
}

// GrandmasterHandler - Handles the /tft/league/v1/grandmaster endpoint
func GrandmasterHandler(w http.ResponseWriter, r *http.Request) {
	// Call the service layer to get data from Riot API
	data, err := service.FetchAllGrandmasters()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response
	sendJSONResponse(w, http.StatusOK, data)
}

// MasterHandler - Handles the /tft/league/v1/master endpoint
func MasterHandler(w http.ResponseWriter, r *http.Request) {
	// Call the service layer to get data from Riot API
	data, err := service.FetchAllMasters()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response
	sendJSONResponse(w, http.StatusOK, data)
}

// CORS Middleware
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://jcb2g11.github.io")    // Allow only your frontend
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allow only necessary methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow necessary headers
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
