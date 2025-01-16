package main

import "fmt"

type House struct {
	balance float64
}

func NewHouse(initialBalance float64) *House {
	return &House{balance: initialBalance}
}

func (h *House) CheckBalance() float64 {
	return h.balance
}

func Main() {
	house := NewHouse(10000.0)
	fmt.Printf("Current house balance: $%.2f\n", house.CheckBalance())
}
