package routes

import (
	internal_websocket "github.com/ZiplEix/super_snake/api/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func webSocketRoutes(app *fiber.App, hub *internal_websocket.Hub) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		hub.HandleNewClient(c)
	}))
}
