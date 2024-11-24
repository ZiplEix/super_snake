package controllers

import (
	"encoding/json"
	"log"

	"github.com/ZiplEix/super_snake/api/hub"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type GameParamsReq struct {
	NbPlayerMax uint `json:"nbPlayerMax"`
	MapHeight   uint `json:"mapHeight"`
	MapWidth    uint `json:"mapWidth"`
}

func CreateGame(c *fiber.Ctx) error {
	gameCreatorID := c.Locals("userId").(uint)

	var req GameParamsReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	gameID := hub.MainHub.CreateGame(gameCreatorID, req.MapHeight, req.MapWidth, req.NbPlayerMax)
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
	MapHeight    uint   `json:"mapHeight"`
	MapWidth     uint   `json:"mapWidth"`
	NbPlayer     uint   `json:"nbPlayer"`
	NbPlayerMax  uint   `json:"nbPlayerMax"`
	GameState    int    `json:"gameState"`
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
		MapHeight:    game.MapHeight,
		MapWidth:     game.MapWidth,
		NbPlayer:     uint(len(game.Players)) + 1,
		NbPlayerMax:  game.NbPlayerMax,
		GameState:    int(game.GameState),
	}
	data, _ := json.Marshal(infos)

	log.Printf("Game %s infos requested", data)

	return c.SendString(string(data))
}
