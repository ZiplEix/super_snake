package websocket

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type TyleType int

const (
	EmptyTyle TyleType = iota
	SnakeTyle
	FoodTyle
)

type Tyle struct {
	Type TyleType
}

type Direction Position

var (
	Up    = Direction{X: 0, Y: -1}
	Down  = Direction{X: 0, Y: 1}
	Left  = Direction{X: -1, Y: 0}
	Right = Direction{X: 1, Y: 0}
)

type Position struct {
	X int
	Y int
}

type Snake struct {
	Body      []Position
	Direction Direction
	Growing   bool
	Color     string
}

type Food struct {
	Position Position
}

type Board struct {
	Width  int
	Height int
	Food   Food
}

type Game struct {
	Snakes map[uint]*Snake
	Board  Board
	colors []string
}

func NewGame(boardWidth, boardHeight, nbMaxSnake int) *Game {
	game := &Game{
		Snakes: make(map[uint]*Snake),
		Board: Board{
			Width:  boardWidth,
			Height: boardHeight,
		},
		colors: GenerateSnakeColors(nbMaxSnake),
	}

	game.spawnFood()

	fmt.Printf("Game created: %v\n", game.String())

	return game
}

func (g *Game) AddSnake(playerID uint) {
	fmt.Printf("Adding snake for player %d\n", playerID)

	// check if player already exist
	if _, ok := g.Snakes[playerID]; ok {
		fmt.Printf("Player %d already has a snake\n", playerID)
		return
	}

	head := Position{
		X: rand.Intn(g.Board.Width),
		Y: rand.Intn(g.Board.Height),
	}
	direction := Down
	g.Snakes[playerID] = &Snake{
		Body: []Position{
			head,
			{
				X: head.X + direction.X,
				Y: head.Y + direction.Y,
			},
			{
				X: head.X + direction.X*2,
				Y: head.Y + direction.Y*2,
			},
		},
		Direction: direction,
		Growing:   false,
		Color:     g.colors[len(g.Snakes)],
	}
}

func (g *Game) spawnFood() {
	g.Board.Food.Position = Position{
		X: rand.Intn(g.Board.Width),
		Y: rand.Intn(g.Board.Height),
	}
}

func (g *Game) MoveSnake(playerID uint) {
	snake := g.Snakes[playerID]
	head := snake.Body[0]
	var newHead Position

	switch snake.Direction {
	case Up:
		newHead = Position{X: head.X, Y: head.Y - 1}
	case Down:
		newHead = Position{X: head.X, Y: head.Y + 1}
	case Left:
		newHead = Position{X: head.X - 1, Y: head.Y}
	case Right:
		newHead = Position{X: head.X + 1, Y: head.Y}
	}

	// Check for collisions
	if newHead.X < 0 || newHead.Y < 0 || newHead.X >= g.Board.Width || newHead.Y >= g.Board.Height {
		// Snake hits the wall, game over logic here
		return
	}

	// Check for self-collision
	for _, part := range snake.Body {
		if part == newHead {
			// Snake collides with itself
			return
		}
	}

	snake.Body = append([]Position{newHead}, snake.Body...)

	if newHead == g.Board.Food.Position {
		snake.Growing = true
		g.spawnFood()
	}

	if !snake.Growing {
		snake.Body = snake.Body[:len(snake.Body)-1]
	} else {
		snake.Growing = false
	}
}

func (g *Game) ChangeSnakeDirection(playerId uint, direction string) {
	// get the snake
	snake, ok := g.Snakes[playerId]
	if !ok {
		fmt.Printf("Snake not found for player %d\n", playerId)
		return
	}

	// check if the direction is valid
	var newDirection Direction
	switch direction {
	case "up":
		newDirection = Up
	case "down":
		newDirection = Down
	case "left":
		newDirection = Left
	case "right":
		newDirection = Right
	default:
		fmt.Printf("Invalid direction: %s\n", direction)
		return
	}

	// check if the new direction is opposite to the current direction
	if newDirection.X == -snake.Direction.X && newDirection.Y == -snake.Direction.Y {
		fmt.Printf("Invalid direction: %s, cannot go in opposite direction\n", direction)
		return
	}

	snake.Direction = newDirection

	fmt.Printf("Snake %d direction changed to %s\n", playerId, direction)
}

func (g *Game) Update() bool {
	for playerID := range g.Snakes {
		g.MoveSnake(playerID)
	}

	return true
}

func (g *Game) String() string {
	return fmt.Sprintf("Snakes: %+v, Food: %+v", g.Snakes, g.Board.Food.Position)
}

type TyleContent struct {
	Type     TyleType `json:"type"`
	PlayerID uint     `json:"playerID"`
	Color    string   `json:"color"`
}

type RawBoardTyle struct {
	X       int         `json:"x"`
	Y       int         `json:"y"`
	Content TyleContent `json:"content"`
}

func (g *Game) GetFullBoardStatus() []byte {
	var rawBoard []RawBoardTyle

	rawBoard = append(rawBoard, RawBoardTyle{
		X: g.Board.Food.Position.X,
		Y: g.Board.Food.Position.Y,
		Content: TyleContent{
			Type: FoodTyle,
		},
	})

	for playerID, snake := range g.Snakes {
		for i, part := range snake.Body {
			color := snake.Color
			if i == 0 {
				darkenedColor, err := DarkenHexColor(color, 30)
				if err != nil {
					fmt.Printf("Error while darkening color: %s\n", err)
				} else {
					color = darkenedColor
				}
			}
			rawBoard = append(rawBoard, RawBoardTyle{
				X: part.X,
				Y: part.Y,
				Content: TyleContent{
					Type:     SnakeTyle,
					PlayerID: playerID,
					Color:    color,
				},
			})
		}
	}

	data, err := json.Marshal(rawBoard)
	if err != nil {
		fmt.Printf("Error while marshalling board data: %s\n", err)
		return nil
	}

	return data
}
