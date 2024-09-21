package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 { // Errors in 400 range is client side and > 499 - server side
		log.Println("Responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"` // Specify ho (Un)Marshall function to convert the struct into JSON
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, playload interface{}) {
	data, err := json.Marshal(playload)

	if err != nil {
		log.Printf("Failed to marshal JSON response: %v\n", playload)
		w.WriteHeader(500) // Internal service error
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code) // Everything went well
	w.Write(data)        // Write JSON data
}
