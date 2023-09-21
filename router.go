package main

import "github.com/gofiber/fiber/v2"

func mainRouter(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from router.go!")
	})
}
