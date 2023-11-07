package pages

import "github.com/gofiber/fiber/v2"

func IndexHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":       "Home",
		"Description": "🚀 A Web Application Template Powered by HTMX + Go Fiber + Tailwind 🚀",
	})
}
