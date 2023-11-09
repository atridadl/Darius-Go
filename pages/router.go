package pages

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	// Page Specific Middleware
	app.Use(jwtware.New(jwtware.Config{
		// CHANGE THIS SECRET! This is used for demo purposes only! You should never hardcode your secret in your code like this!
		SigningKey: jwtware.SigningKey{Key: []byte("CHANGEME")},
		Filter: func(c *fiber.Ctx) bool {
			return c.Path() != "/restricted"
		},
		TokenLookup: "cookie:token",
	}))

	// Page Routes
	app.Get("/", IndexHandler)
	app.Get("/restricted", Restricted)
}
