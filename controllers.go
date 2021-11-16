package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type short struct {
	Url      string `json:"url"`
	ShortUrl string `json:"shortUrl"`
}

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint64(len(alphabet))
)

var userCollection = db().Database("shortenedUrl").Collection("urls")

func createShortenedUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // for adding       //Content-type
	var shortModel short
	shortModel.ShortUrl = "https://my-url/" + getShortUrl(rand.Uint64())

	err := json.NewDecoder(r.Body).Decode(&shortModel) // storing in shortModel   //variable of type user
	if err != nil {
		fmt.Print(err)
	}
	var result primitive.M
	errV := userCollection.FindOne(context.TODO(), bson.M{"url": shortModel.Url}).Decode(&result)
	if errV != nil {
		insertResult, err := userCollection.InsertOne(context.TODO(), shortModel)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(insertResult.InsertedID) //  return resulted ID
	}
	json.NewEncoder(w).Encode(result) // If Url Exist return Result
}

func getShortUrl(number uint64) string {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(11)

	for ; number > 0; number = number / length {
		encodedBuilder.WriteByte(alphabet[(number % length)])
	}

	return encodedBuilder.String()
}

func getUrl(w http.ResponseWriter, r *http.Request) {
}
