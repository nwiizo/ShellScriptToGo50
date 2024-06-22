package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/users", func(c *fiber.Ctx) error {
		users := []map[string]interface{}{
			{"id": 1, "name": "John Doe", "email": "john@example.com"},
			{"id": 2, "name": "Jane Smith", "email": "jane@example.com"},
		}
		return c.JSON(users)
	})

	log.Fatal(app.Listen(":3000"))
}
