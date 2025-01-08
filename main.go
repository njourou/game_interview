package main

import (
    "log"
    "net/http"
    "time"
    "math/rand"
)

func main() {
    rand.Seed(time.Now().UnixNano())
//start the server
    http.HandleFunc("/spin", handleSpin)
    log.Println("Starting slot machine server on :8080")
	
    log.Fatal(http.ListenAndServe(":8080", nil))
}