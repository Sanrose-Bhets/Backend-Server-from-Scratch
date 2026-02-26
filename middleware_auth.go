package main

import (
	"fmt"
	"net/http"

	"github.com/Sanrose-Bhets/Backend-Server-from-Scratch/internal/auth"
	"github.com/Sanrose-Bhets/Backend-Server-from-Scratch/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

//work, return a closure, a new annon function is returned with the same signature as a http handler function
//diff being, the difference being, shall have access to everything within the API config so we can query the database

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//this is a aunthenticated endpoint
		//get own user info, give us api key, not the only end point tho
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("ACouldn't get user: %v", err))
			return
		}
		handler(w, r, user)
	}

}
