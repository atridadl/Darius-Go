package main

import (
	"darius/api"
	"darius/pages"
	"log"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/template/html/v2"
)

//go:generate npx tailwindcss -i ./main.css -o ./public/styles.css

func main() {
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "base",
	})

	// Middleware
	app.Use(jwtware.New(jwtware.Config{
		// CHANGE THIS SECRET! This is used for demo purposes only! You should never hardcode your secret in your code like this!
		SigningKey: jwtware.SigningKey{Key: []byte("CHANGEME")},
		Filter: func(c *fiber.Ctx) bool {
			return c.Path() != "/restricted"
		}}))
	app.Use(etag.New())

	// Static Routes
	app.Static("/", "./public", fiber.Static{
		Compress: true,
	})

	// Page Routes
	pages.Register(app)

	// API Routes
	api.Register(app)

	log.Fatal(app.Listen(":3000"))
}
