package routes

import (
	"github.com/ZiplEix/super_snake/api/controllers"
	"github.com/ZiplEix/super_snake/api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func webSocketRoutes(app *fiber.App) {
	app.Use("/ws", middleware.Protected, func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// Route pour créer une partie
	app.Get("/game/create", middleware.Protected, controllers.CreateGame)

	// Route WebSocket pour rejoindre une partie
	app.Get("/ws/:id", middleware.Protected, websocket.New(controllers.JoinGame))
}
