package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(
		fiber.Config{
			AppName:       "v-lock API Server (Golang)",
			CaseSensitive: true,
			ServerHeader:  "v-lock API Server (Golang)",
		})

	app.Static("/", "./public")

	var prefix string = "/api"
	api := app.Group(prefix)

	mainRouter(api)
	app.Listen(":3000")
}
