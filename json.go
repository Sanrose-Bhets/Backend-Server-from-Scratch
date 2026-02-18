package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// a function that sends a error message if something were to go wrong.
// similar to the respondWithJson function, with a change in its last parameter,
// this function will format that message into a consistent JSON object each time
func respondWithError(w http.ResponseWriter, code int, msg string) {
	//code till 499 means,(client Side Errors) someone using the API funcky way, aka users faults/errors
	//code in the 500 means its cause by the server, thus our fault
	if code > 499 {
		log.Println("Responding with a 5xx error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Fail to Marshal the JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	//adds a response header to the http request, saying:= we're responding with a content type of application/json
	//which is the standard for json responses
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
