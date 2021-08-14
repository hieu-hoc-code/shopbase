package main

import (
	"github.com/gofiber/fiber"
	"./routes"
)

func main() {
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}