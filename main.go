package main

import (
	"github.com/gofiber/fiber/v2"
)

type todos struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func fiberGet() {
}

func main() {
	app := fiber.New()

	app.Get("/", TodoServer)

	app.Listen(":8080")
}

// TodoServer Function
func TodoServer(c *fiber.Ctx) error {
	return c.SendString("Hello, world!")
}
