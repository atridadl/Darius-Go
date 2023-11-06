package api

import "github.com/gofiber/fiber/v2"

func HelloHandler(c *fiber.Ctx) error {
	return c.SendString(`<p id="hello">Hello! This came from the server!</p>`)
}
