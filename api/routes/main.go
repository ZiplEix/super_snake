package routes

import (
	"github.com/ZiplEix/super_snake/api/websocket"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, hub *websocket.Hub) {
	authRoutes(app)
	webSocketRoutes(app, hub)
}
