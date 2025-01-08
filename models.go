package main

type PlayRequest struct {
	BetAmount float64 `json:"bet_amount"`
}

type PlayResponse struct {
	Outcome    string  `json:"outcome"` // win or loss required
	WinAmount  float64 `json:"win_amount"`
	FreeGames  int     `json:"free_games"`
	Message    string  `json:"message"`
}
