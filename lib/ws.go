package lib

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type client struct{}

var clients = make(map[*websocket.Conn]client)
var register = make(chan *websocket.Conn)
var broadcast = make(chan string)
var unregister = make(chan *websocket.Conn)

func RunHub() {
	for {
		// Handle wither a message is received or a client is registered/unregistered
		select {
		case connection := <-register:
			clients[connection] = client{}
			log.Println("connection registered")

		case message := <-broadcast:
			log.Println("message received:", message)

			// Send the message to all clients
			for connection := range clients {
				if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					log.Println("write error:", err)

					unregister <- connection
					connection.WriteMessage(websocket.CloseMessage, []byte{})
					connection.Close()
				}
			}

		case connection := <-unregister:
			// Remove the client from the hub
			delete(clients, connection)

			log.Println("connection unregistered")
		}
	}
}

func WsUseHandler(wsc *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(wsc) {
		wsc.Locals("allowed", true)
		return wsc.Next()
	}
	return fiber.ErrUpgradeRequired
}

func WsGetHandler(c *websocket.Conn) {
	// When the function returns, unregister the client and close the connection
	defer func() {
		unregister <- c
		c.Close()
	}()

	// Register the client
	register <- c

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("read error:", err)
			}
			// Close the connection with an error status
			return
		}

		if messageType == websocket.TextMessage {
			// Broadcast the received message
			broadcast <- string(message)
		} else {
			log.Println("websocket message received of type", messageType)
		}
	}
}
