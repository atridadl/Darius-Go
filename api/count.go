package api

import (
	"fmt"
	"log"
	"net/url"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
)

var count = 0

func GetCounthandler(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf(`<p hx-id="countfromserver" id="countfromserver">%d</p>`, count))
}

func DecrementCountHandler(c *fiber.Ctx) error {
	count--
	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws"}
	wsc, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Printf("handshake failed with status %d", resp.StatusCode)
		log.Fatal("dial:", err)
	}
	wsc.WriteMessage(websocket.TextMessage, []byte(string(fmt.Sprintf(`<p hx-id="countfromserver" id="countfromserver">%d</p>`, count))))

	defer wsc.Close()
	return err
}

func IncrementCountHandler(c *fiber.Ctx) error {
	count++
	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws"}
	wsc, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Printf("handshake failed with status %d", resp.StatusCode)
		log.Fatal("dial:", err)
	}
	wsc.WriteMessage(websocket.TextMessage, []byte(string(fmt.Sprintf(`<p hx-id="countfromserver" id="countfromserver">%d</p>`, count))))

	defer wsc.Close()

	return err
}
