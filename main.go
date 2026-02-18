package main

//A JSON Rest API Backend Server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello world")

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT was not found in the environment")
	}

	fmt.Println("Port: ", portString)

	//creating a new router object with chi router
	router := chi.NewRouter()

	//enables users to send requests from any browsers
	//passing cors configurations
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//now hooking up multiple the handlers , using the chi router, to a specific http method and path
	//creating a new router to Mount it
	v1Router := chi.NewRouter()

	//changed from HandleFunc to get to make it only fire in get requests
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	//So nesting a V1 router under the /v1 path, and hooking up the readiness function to the /ready path
	//so the full path for this request will be /v1/ready
	//this has been done so that if there are any breaking chnages in the  future, we have 2 handlers, one v1 and other v2, a standard practoce
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server is Starting on Port: %v", portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
