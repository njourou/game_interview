package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	//seed the random number generator

	rand.Seed(time.Now().UnixNano())
	//start the server
	http.HandleFunc("/spin", handleSpin)

	log.Println("starting the prrogram on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
