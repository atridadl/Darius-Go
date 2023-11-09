package api

import (
	"darius/lib"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	// API Specific Middleware
	app.Use(jwtware.New(jwtware.Config{
		// CHANGE THIS SECRET! This is used for demo purposes only! You should never hardcode your secret in your code like this!
		SigningKey: jwtware.SigningKey{Key: []byte("CHANGEME")},
		Filter: func(c *fiber.Ctx) bool {
			return c.Path() != "/api/restricted"
		},
	}))

	app.Use(jwtware.New(jwtware.Config{
		// CHANGE THIS SECRET! This is used for demo purposes only! You should never hardcode your secret in your code like this!
		SigningKey: jwtware.SigningKey{Key: []byte("CHANGEME")},
		Filter: func(c *fiber.Ctx) bool {
			return c.Path() != "/api/loginform"
		},
		TokenLookup: "cookie:token",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if c.Path() == "/api/loginform" {
				return c.Next()
			} else {
				return c.SendString("Missing or malformed JWT")
			}
		},
	}))

	// Register the API routes
	app.Get("/api/hello", HelloHandler)
	app.Get("/api/countplus", IncrementCountHandler)
	app.Get("/api/countminus", DecrementCountHandler)
	app.Get("/api/loginform", LoginFormHandler)

	app.Get("/api/count", GetCounthandler)

	// Register Auth API routes
	app.Get("/api/restricted", RestrictedHandler)
	app.Post("/api/token", lib.GenerateJWTHandler)

	// Register the websocket routes
	go lib.RunHub()
	app.Use("/ws", lib.WsUseHandler)
	app.Get("/ws", websocket.New(lib.WsGetHandler))
}
