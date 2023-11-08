package api

import (
	"darius/lib"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	// Register the API routes
	app.Get("/hello", HelloHandler)
	app.Get("/countplus", IncrementCountHandler)
	app.Get("/countminus", DecrementCountHandler)
	app.Get("/count", GetCounthandler)

	// Private route
	app.Get("/restricted", lib.Restricted)
	app.Post("/token/login", lib.GetJWT)

	// Register the websocket routes
	go lib.RunHub()
	app.Use("/ws", lib.WsUseHandler)
	app.Get("/ws", websocket.New(lib.WsGetHandler))
}
