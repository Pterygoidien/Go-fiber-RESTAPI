package main

import "github.com/gofiber/fiber/v2"

func mainRouter(api fiber.Router) {
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from router.go!")
	})

	api.Group("/auth", authRouter)

}
