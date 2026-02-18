package main

import "net/http"

//have to use this speicfic function signature to define a http handler where the go standard library expects

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
