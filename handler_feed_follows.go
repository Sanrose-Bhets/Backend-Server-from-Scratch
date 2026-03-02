package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sanrose-Bhets/Backend-Server-from-Scratch/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// have to use this speicfic function signature to define a http handler where the go standard library expects
// as you cant change the function signature, but still wish to enter more parameters,
// we change it to a method
func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	//this handler needs to take a input a JSON body
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error with JSON:%v", err))
		return
	}

	FeedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Create feed Follow: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowtoFeedFollow(FeedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	//this handler needs to take a input a JSON body

	feedFollows, err := apiCfg.DB.GetFeedFollow(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Get feed Follows: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowstoFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	//this handler needs to take a input a JSON body
	feedfollowIdstr := chi.URLParam(r, "feedFollowID")
	feedfollowid, err := uuid.Parse(feedfollowIdstr)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Parse feed follow id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedfollowid,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Delete the Feed the user Followed: %v", err))
		return
	}
	respondWithJSON(w, 201, struct{}{})
}
