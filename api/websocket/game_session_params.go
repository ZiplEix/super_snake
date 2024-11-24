package websocket

type Params struct {
	GameLeaderID uint
	NbPlayerMax  uint
	MapHeight    uint
	MapWidth     uint
	GameState    State
}

func NewParams(gameLeaderID, nbPlayerMax, mapHeight, mapWidth uint) Params {
	return Params{
		GameLeaderID: gameLeaderID,
		NbPlayerMax:  nbPlayerMax,
		MapHeight:    mapHeight,
		MapWidth:     mapWidth,
		GameState:    WaitingToStart,
	}
}
