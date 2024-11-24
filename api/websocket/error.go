package websocket

import "errors"

var (
	ErrGameNotFound = errors.New("game not found")
	ErrGameFull     = errors.New("game is full")
)
