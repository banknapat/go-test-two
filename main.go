package main

import (
	"github.com/banknapat/go-test-two/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Routes(app)

	app.Listen(":3000")
}
