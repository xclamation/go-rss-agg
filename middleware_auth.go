package main

import (
	"fmt"
	"net/http"

	"github.com/xclamation/go-rss-agg/internal/auth"
	"github.com/xclamation/go-rss-agg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // Use a clousure to use user's information but return func that satisfies http.HandlerFunc structure
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err)) // 403 - permisson errors code
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey) // Context() allows you to track smth that happens across multiple goroutines
		// and *IMPORTANT* you can cancel it - effectively kill http request - when it necessary
		// used in http hanlders

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err)) // 403 - permisson errors code
			return
		}

		handler(w, r, user)
	}
}
