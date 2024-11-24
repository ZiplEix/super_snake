package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
)

type Client struct {
	id   uint
	game *GameSession
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) ReadPump() {
	defer func() {
		c.game.Unregister <- c
		_ = c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Message reçu du client (longueur %d): %s\n", len(message), string(message))
			}
			break
		}

		log.Printf("Message brut reçu du client: %s\n", string(message))

		// c.game.Broadcast <- message

		var evt Event
		if err := json.Unmarshal(message, &evt); err != nil {
			log.Printf("Erreur lors du décodage du message JSON: %s\n", err)
			continue
		}

		switch evt.Type {
		case "game_control":
			fmt.Println("Game control event received")
			var controlEvt GameControlEvent
			if err := json.Unmarshal(evt.Data, &controlEvt); err != nil {
				log.Printf("Erreur lors du décodage de l'événement de contrôle du jeu: %s\n", err)
				continue
			}
			c.game.HandleGameControlEvent(controlEvt)
			fmt.Println("Game control event handled")
		default:
			log.Printf("Type d'événement inconnu: %s\n", evt.Type)
		}

		fmt.Println("Event received:", evt)
	}
}

func (c *Client) WritePump() {
	defer c.conn.Close()
	for message := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			break
		}
	}
}
