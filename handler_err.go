package main

import "net/http"

// Checking that server is alive and running
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong") // 400 - client error
}