package http

import (
	"encoding/json"
	"net/http"

	"github.com/hamed-amini-dev/stripe-go-scripts/internal/response"
)

/*
Set Browser origin
*/
func setOrigins(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

/*
respondJSON- Returns the response in JSON format to the response
writer | Adds content type and charset to the header map | Sends an
HTTP response header with the provided status code.
*/
func respondJSON(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
	w.WriteHeader(status)
}

// respondError- Returns an error as a response
func respondError(w http.ResponseWriter, status int, msg string) {
	errBody := response.GenericResponse{
		Messages: []string{msg},
	}
	respondJSON(w, status, errBody)
}
