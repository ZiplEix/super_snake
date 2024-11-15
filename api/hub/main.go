package hub

import (
	"log"
	"sync"

	"github.com/ZiplEix/super_snake/api/utils"
	internal_websocket "github.com/ZiplEix/super_snake/api/websocket"
)

var (
	MainHub *Hub
)

type Hub struct {
	games map[string]*internal_websocket.Game
	mu    sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		games: make(map[string]*internal_websocket.Game),
	}
}

func (h *Hub) CreateGame() string {
	h.mu.Lock()
	defer h.mu.Unlock()

	gameID := utils.GenerateGameID()
	game := internal_websocket.NewGame(gameID)
	h.games[gameID] = game

	go game.Run()

	log.Printf("Game %s created", gameID)
	return gameID
}

func (h *Hub) GetGame(gameID string) *internal_websocket.Game {
	h.mu.Lock()
	defer h.mu.Unlock()

	game, exists := h.games[gameID]
	if !exists {
		return nil
	}
	return game
}

func (h *Hub) JoinGame(gameID string, client *internal_websocket.Client) bool {
	game, exists := h.games[gameID]
	if !exists || len(game.Players) >= 2 {
		return false
	}
	game.Register <- client
	return true
}
