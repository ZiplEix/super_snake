package websocket

// type Hub struct {
// 	games map[string]*Game
// 	mu    sync.Mutex
// }

// func NewHub() *Hub {
// 	return &Hub{
// 		games: make(map[string]*Game),
// 	}
// }

// func (h *Hub) CreateGame() string {
// 	h.mu.Lock()
// 	defer h.mu.Unlock()

// 	gameID := utils.GenerateGameID()
// 	game := NewGame(gameID)
// 	h.games[gameID] = game

// 	go game.Run()

// 	log.Printf("Game %s created", gameID)
// 	return gameID
// }

// func (h *Hub) GetGame(gameID string) *Game {
// 	h.mu.Lock()
// 	defer h.mu.Unlock()

// 	game, exists := h.games[gameID]
// 	if !exists {
// 		return nil
// 	}
// 	return game
// }

// func (h *Hub) JoinGame(gameID string, client *Client) bool {
// 	game, exists := h.games[gameID]
// 	if !exists || len(game.Players) >= 2 {
// 		return false
// 	}
// 	game.Register <- client
// 	return true
// }
