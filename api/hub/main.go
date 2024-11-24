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
	games map[string]*internal_websocket.GameSession
	mu    sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		games: make(map[string]*internal_websocket.GameSession),
		mu:    sync.Mutex{},
	}
}

func (h *Hub) CreateGame(gameCreatorID, mapHeight, mapWidth, nbPlayerMax uint) string {
	h.mu.Lock()
	defer h.mu.Unlock()

	gameID := utils.GenerateGameID()

	params := internal_websocket.NewParams(gameCreatorID, nbPlayerMax, mapHeight, mapWidth)

	game := internal_websocket.NewGameSession(gameID, params)
	h.games[gameID] = game

	go game.Run()

	log.Printf("Game %s created, with params: %v", gameID, params)
	return gameID
}

func (h *Hub) GetGame(gameID string) *internal_websocket.GameSession {
	h.mu.Lock()
	defer h.mu.Unlock()

	game, exists := h.games[gameID]
	if !exists {
		return nil
	}
	return game
}

func (h *Hub) JoinGame(gameID string, client *internal_websocket.Client) error {
	game, exists := h.games[gameID]
	if !exists {
		return internal_websocket.ErrGameNotFound
	}

	if len(game.Players) >= int(game.NbPlayerMax) {
		return internal_websocket.ErrGameFull
	}

	game.Register <- client
	return nil
}
