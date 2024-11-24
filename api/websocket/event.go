package websocket

import "encoding/json"

type Event struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func (e Event) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

type GameControlEvent struct {
	Action string `json:"action"`
}

type PlayerMoveEvent struct {
	Direction string `json:"direction"`
}
