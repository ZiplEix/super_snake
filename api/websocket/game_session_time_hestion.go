package websocket

import "time"

type TimeGestion struct {
	GameTicker             *time.Ticker
	InactivityTimer        *time.Timer
	InactivityChan         <-chan time.Time
	PlayerConnectionTicker *time.Ticker
}

func NewTimeGestion() *TimeGestion {
	return &TimeGestion{
		GameTicker:             time.NewTicker(300 * time.Millisecond),
		InactivityTimer:        time.NewTimer(5 * time.Minute),
		InactivityChan:         make(<-chan time.Time),
		PlayerConnectionTicker: time.NewTicker(5 * time.Second),
	}
}
