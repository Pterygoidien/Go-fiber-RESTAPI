package main

import "github.com/gofiber/fiber/v2"

func authRouter(router fiber.Router) {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from authRouter!")
	})

}
