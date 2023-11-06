package main

import (
	"darius/api"
	"darius/pages"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

//go:generate npx tailwindcss -i ./main.css -o ./public/styles.css

func main() {
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Static Routes
	app.Static("/", "./public")

	// Page Routes
	app.Get("/", pages.IndexHandler)

	// API Routes
	app.Get("/hello", api.HelloHandler)

	log.Fatal(app.Listen(":3000"))
}
