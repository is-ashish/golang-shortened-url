package main

import (
	"net/http"
)

type short struct {
	Url      string `json:"url"`
	ShortUrl string `json:"shortUrl"`
}

var userCollection = db().Database("shortenedUrl").Collection("urls")

func createShortenedUrl(w http.ResponseWriter, r *http.Request) {
}

func getUrl(w http.ResponseWriter, r *http.Request) {
}
