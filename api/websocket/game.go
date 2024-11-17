package websocket

import (
	"log"
	"sync"
	"time"

	"github.com/gofiber/websocket/v2"
)

type Params struct {
	GameLeaderID uint
}

func NewParams(gameLeaderID uint) Params {
	return Params{
		GameLeaderID: gameLeaderID,
	}
}

type Game struct {
	ID         string
	Players    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
	mu         sync.Mutex

	Params
}

func NewGame(id string, gameLeaderID uint) *Game {
	return &Game{
		ID:         id,
		Players:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),

		Params: NewParams(gameLeaderID),
	}
}

func (g *Game) HandleNewClient(conn *websocket.Conn) {
	client := &Client{
		game: g,
		conn: conn,
		send: make(chan []byte, 256),
	}
	g.Register <- client
	go client.ReadPump()
	client.WritePump()
}

func (g *Game) Run() {
	gameTicker := time.NewTicker(50 * time.Millisecond)

	inactivityTimer := time.NewTimer(5 * time.Minute)
	inactivityChan := make(<-chan time.Time)

	playerConnectionTocker := time.NewTicker(5 * time.Second)

	for {
		select {
		case client := <-g.Register:
			g.mu.Lock()

			g.Players[client] = true
			log.Printf("Client connected to game '%s'\n", g.ID)

			if inactivityTimer != nil {
				inactivityTimer.Stop()
				inactivityTimer = nil
				inactivityChan = make(<-chan time.Time)
				log.Printf("Game '%s' has players, inactivity timer stopped\n", g.ID)
			}

			g.mu.Unlock()

		case client := <-g.Unregister:
			if _, ok := g.Players[client]; ok {
				delete(g.Players, client)
				close(client.send)
				log.Printf("Client disconnected from game '%s'\n", g.ID)
			}

			if len(g.Players) == 0 && inactivityTimer == nil {
				inactivityTimer = time.NewTimer(5 * time.Minute)
				inactivityChan = inactivityTimer.C
				log.Printf("Game '%s' has no players, inactivity timer started\n", g.ID)
			}

		case message := <-g.Broadcast:
			for client := range g.Players {
				select {
				case client.send <- message:
					log.Printf("Message sent to player in game %s: %s\n", g.ID, string(message))
				default:
					close(client.send)
					delete(g.Players, client)
				}
			}

		case <-gameTicker.C:
			g.mu.Lock()
			// send hello world to all player
			// for client := range g.Players {
			// 	select {
			// 	case client.send <- []byte("Hello, World!"):
			// 		// log.Printf("Message sent to player in game %s: Hello, World!", g.ID)
			// 		continue
			// 	default:
			// 		close(client.send)
			// 		delete(g.Players, client)
			// 	}
			// }
			g.mu.Unlock()

		case <-playerConnectionTocker.C:
			g.mu.Lock()

			for client := range g.Players {
				// send the message "ping" to all players
				select {
				case client.send <- []byte("ping"):
					log.Printf("Message sent to player in game %s: ping\n", g.ID)
					continue
				default:
					log.Printf("Client disconnected from game '%s'\n", g.ID)
					close(client.send)
					delete(g.Players, client)
				}
			}

			g.mu.Unlock()

		case <-inactivityChan:
			g.mu.Lock()
			if len(g.Players) == 0 {
				log.Printf("Game '%s' has been inactive for 5 minutes, closing game\n", g.ID)
				g.mu.Unlock()
				return
			}

			inactivityTimer.Stop()
			g.mu.Unlock()
			return
		}

	}
}
