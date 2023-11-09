package main

import (
	"darius/api"
	"darius/pages"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/template/html/v2"
)

//go:generate npx tailwindcss -i ./main.css -o ./public/styles.css

func main() {
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Primary Middleware
	app.Use(etag.New())
	app.Use(helmet.New())

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
