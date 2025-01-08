package main

import (
    "encoding/json"
    "net/http"
)

func handleSpin(w http.ResponseWriter, r *http.Request) {

	
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    var requestData struct {
        BetAmount float64 `json:"betAmount"`
    }

    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, "Invalid request ", http.StatusBadRequest)
        return
    }

    result := spin(requestData.BetAmount)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}