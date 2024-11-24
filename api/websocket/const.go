package websocket

type State int

const (
	WaitingToStart State = iota
	Started
	Paused
	Resumed
	GameOver
)
