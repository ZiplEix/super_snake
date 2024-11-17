package controllers

import (
	"encoding/json"
	"log"

	"github.com/ZiplEix/super_snake/api/hub"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func CreateGame(c *fiber.Ctx) error {
	gameCreatorID := c.Locals("userId").(uint)

	gameID := hub.MainHub.CreateGame(gameCreatorID)
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

type GameInfos struct {
	ID           string `json:"id"`
	GameLeaderID uint   `json:"gameLeaderID"`
}

func GetGameInfos(c *fiber.Ctx) error {
	gameID := c.Params("id")
	game := hub.MainHub.GetGame(gameID)
	if game == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Game not found"})
	}

	infos := GameInfos{
		ID:           game.ID,
		GameLeaderID: game.GameLeaderID,
	}
	data, _ := json.Marshal(infos)

	log.Printf("Game %s infos requested", data)

	return c.SendString(string(data))
}
