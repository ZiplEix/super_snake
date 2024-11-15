package controllers

import (
	"github.com/ZiplEix/super_snake/api/hub"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func CreateGame(c *fiber.Ctx) error {
	gameID := hub.MainHub.CreateGame()
	return c.JSON(fiber.Map{"gameID": gameID})
}

func JoinGame(c *websocket.Conn) {
	gameID := c.Params("id")
	game := hub.MainHub.GetGame(gameID)
	if game == nil {
		_ = c.Close()
		return
	}
	game.HandleNewClient(c)
}
