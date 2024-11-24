package websocket

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gofiber/websocket/v2"
)

type GameSession struct {
	ID         string
	Players    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
	Mu         sync.Mutex

	Params
	*Game
	*TimeGestion
}

func NewGameSession(id string, params Params) *GameSession {
	return &GameSession{
		ID:         id,
		Players:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte, 100),

		Params:      params,
		Game:        NewGame(int(params.MapWidth), int(params.MapHeight), int(params.NbPlayerMax)),
		TimeGestion: NewTimeGestion(),
	}
}

func (g *GameSession) HandleNewClient(conn *websocket.Conn) {
	client := &Client{
		id:   conn.Locals("userId").(uint),
		game: g,
		conn: conn,
		send: make(chan []byte, 256),
	}
	g.Register <- client
	go client.ReadPump()
	client.WritePump()
}

func (gs *GameSession) RegisterClient(client *Client) {
	gs.Players[client] = true
	log.Printf("Client connected to game '%s'\n", gs.ID)

	if gs.InactivityTimer != nil {
		gs.InactivityTimer.Stop()
		gs.InactivityTimer = nil
		gs.InactivityChan = make(<-chan time.Time)
		log.Printf("Game '%s' has players, inactivity timer stopped\n", gs.ID)
	}

	gs.Game.AddSnake(client.id)

	boardData := gs.Game.GetFullBoardStatus()

	emitedEvent := Event{
		Type: "StartBoard",
		Data: boardData,
	}

	data, err := emitedEvent.Marshal()
	if err != nil {
		log.Printf("Error while marshalling event: %s\n", err)
		return
	}

	gs.Broadcast <- data
}

func (gs *GameSession) UnregisterClient(client *Client) {
	if _, ok := gs.Players[client]; ok {
		delete(gs.Players, client)
		close(client.send)
		log.Printf("Client disconnected from game '%s'\n", gs.ID)
	}

	if len(gs.Players) == 0 && gs.InactivityTimer == nil {
		gs.InactivityTimer = time.NewTimer(5 * time.Minute)
		gs.InactivityChan = gs.InactivityTimer.C
		log.Printf("Game '%s' has no players, inactivity timer started\n", gs.ID)
	}
}

func (gs *GameSession) BroadcastMessage(message []byte) {
	for client := range gs.Players {
		select {
		case client.send <- message:
			log.Printf("(BROADCAST) Message sent to player in game %s: %s\n", gs.ID, string(message))
		default:
			log.Printf("(BROADCAST) Client channel is full or closed. Removing client from game '%s'\n", gs.ID)
			close(client.send)
			delete(gs.Players, client)
		}
	}
}

func (gs *GameSession) CheckClientHealth() {
	for client := range gs.Players {
		select {
		case client.send <- []byte("ping"):
			// log.Printf("Message sent to player in game %s: ping\n", gs.ID)
			continue
		default:
			log.Printf("Client disconnected from game '%s'\n", gs.ID)
			close(client.send)
			delete(gs.Players, client)
		}
	}
}

func (gs *GameSession) CloseGameForInactivity() bool {
	if len(gs.Players) == 0 {
		log.Printf("Game '%s' has been inactive for 5 minutes, closing game\n", gs.ID)
		return true
	}

	gs.InactivityTimer.Stop()
	return false
}

func (gs *GameSession) Run() {
	for {
		select {
		case client := <-gs.Register:
			gs.RegisterClient(client)

		case client := <-gs.Unregister:
			gs.UnregisterClient(client)

		case message := <-gs.Broadcast:
			gs.BroadcastMessage(message)

		case <-gs.GameTicker.C:
			gs.Update()

		case <-gs.PlayerConnectionTicker.C:
			gs.CheckClientHealth()

		case <-gs.InactivityChan:
			if gs.CloseGameForInactivity() {
				return
			}
		}

	}
}

func (gs *GameSession) HandleGameControlEvent(evt GameControlEvent) {
	switch evt.Action {
	case "start":
		gs.GameState = Started
	case "pause":
		gs.GameState = Paused
	case "resume":
		gs.GameState = Resumed
	default:
		log.Printf("Unknown game control event: %s\n", evt.Action)
	}

	fmt.Println("Game state:", gs.GameState)

	emitedEvent := Event{
		Type: "GameStart",
	}

	data, err := emitedEvent.Marshal()
	if err != nil {
		log.Printf("Error while marshalling event: %s\n", err)
		return
	}

	fmt.Println("Broadcasting game start event")

	gs.Broadcast <- data
}
