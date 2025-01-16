package main

import (
	"math/rand"
)

type Symbol struct {
	Name       string
	Weight     int
	Multiplier float64
}

// just symbols for testing
var symbols = []Symbol{
	{"🍒", 35, 2},
	{"🍋", 30, 2.5},
	{"🍊", 25, 3},
	{"⭐", 15, 5},
	{"💎", 10, 10},
	{"7️⃣", 5, 20},
}

func getRandomSymbol() string {
	totalWeight := 0
	for _, symbol := range symbols {
		totalWeight += symbol.Weight
	}

	r := rand.Intn(totalWeight)

	current := 0
	for _, symbol := range symbols {
		current += symbol.Weight
		if r < current {
			return symbol.Name
		}
	}

	return symbols[0].Name // fallback
}

func getMultiplier(symbolName string) float64 {
	for _, symbol := range symbols {
		if symbol.Name == symbolName {
			return symbol.Multiplier
		}
	}
	return 1.0
}
