package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path
	//Routes
	s.HandleFunc("/createShortenedUrl", createShortenedUrl).Methods("POST")
	s.HandleFunc("/getUrl", getUrl).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
