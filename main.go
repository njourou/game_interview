package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func PlayHandler(c *fiber.Ctx) error {
	apiKey := c.Get("Authorization")
	validApiKey := os.Getenv("API_KEY")

	if apiKey != "Bearer "+validApiKey {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Invalid API Key",
		})
	}

	return c.JSON(fiber.Map{
		"message": "API Key is valid. Proceeding with request.",
	})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	app := fiber.New()
	app.Post("/play", PlayHandler)

	port := ":3000"
	log.Printf("Server is running on port %s", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
