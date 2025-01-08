package main

import (


	"github.com/gofiber/fiber/v2"
)


func PlayHandler(c *fiber.Ctx) error {

	request := new(PlayRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload. Please provide a valid bet amount.",
		})
	}

	if request.BetAmount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bet amount must be greater than zero.",
		})
	}

	
	outcome := simulateGame(request.BetAmount)

	
	return c.JSON(outcome)
}
