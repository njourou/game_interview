package main

import "math/rand"

// simulate the game
func simulateGame(betAmount float64) PlayResponse {
	outcome := "loss"
	winAmount := 0.0
	freeGames := 0
	message := "Unfortunately, you didn’t win this time. Try again!"

	// 50% probability
	if rand.Intn(2) == 0 {
		outcome = "win"
		winAmount = betAmount * float64(rand.Intn(10)+1) //  multiplier between 1 and 10
		message = "Congratulations! You hit a win."
	}

	// Add a chance for free games 
	if rand.Intn(100) < 25 {
		freeGames = rand.Intn(5) + 3 
		message += " Plus, you earned some free games!"
	}

	return PlayResponse{
		Outcome:   outcome,
		WinAmount: winAmount,
		FreeGames: freeGames,
		Message:   message,
	}
}
