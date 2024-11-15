package websocket

import "github.com/gofiber/websocket/v2"

type Client struct {
	game *Game
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
			break
		}
		c.game.Broadcast <- message
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
