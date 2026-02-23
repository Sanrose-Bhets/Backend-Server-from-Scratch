package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sanrose-Bhets/Backend-Server-from-Scratch/internal/database"
	"github.com/google/uuid"
)

// have to use this speicfic function signature to define a http handler where the go standard library expects
// as you cant change the function signature, but still wish to enter more parameters,
// we change it to a method
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	//this handler needs to take a input a JSON body
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error with JSON:%v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Create user: %v", err))
		return
	}

	respondWithJSON(w, 200, user)
}
