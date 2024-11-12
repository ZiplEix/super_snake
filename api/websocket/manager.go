package websocket

import "github.com/gofiber/websocket/v2"

func (h *Hub) HandleNewClient(conn *websocket.Conn) {
	client := &Client{
		hub:  h,
		conn: conn,
		send: make(chan []byte, 256),
	}
	h.register <- client

	go client.WritePump()
	client.ReadPump()
}
