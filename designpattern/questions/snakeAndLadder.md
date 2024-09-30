```go

package models

import (
	"fmt"
	"math/rand"
)

type Board struct {
	Snakes  []Snake
	Ladders []Ladder
}

func (b Board) IsSankePresent(pos int) (bool, int) {
	for i := 0; i < len(b.Snakes); i++ {
		if pos == b.Snakes[i].Start {
			return true, b.Snakes[i].End
		}
	}
	return false, 0
}
func (b Board) IsLadderPresent(pos int) (bool, int) {
	for i := 0; i < len(b.Ladders); i++ {
		if pos == b.Ladders[i].Start {
			return true, b.Ladders[i].End
		}
	}
	return false, 0
}

type Snake struct {
	Start int
	End   int
}

type Ladder struct {
	Start int
	End   int
}

type Dice struct {
	Value     int
	MaxNumber int
}

func NewDice(val int) Dice {
	return Dice{Value: val, MaxNumber: 6}
}

func (d Dice) GenerateNumber() int {
	return d.Value * (1 + rand.Intn(d.MaxNumber))
}

type Game struct {
	Board       Board
	Players     []*PlayerPos
	CurrentTurn *PlayerPos
	Dice        Dice
}

func NewGame(Board Board, Player []Player, Dice Dice) *Game {
	players := make([]*PlayerPos, len(Player))
	for i := 0; i < len(Player); i++ {
		players[i] = NewPlayerPos(Player[i])
	}
	return &Game{
		Board:       Board,
		Players:     players,
		CurrentTurn: players[0],
		Dice:        Dice,
	}
}

func (g *Game) PlayGame() {
	len := len(g.Players)
	i := 0
	for {
		fmt.Printf("Player Turn %s%d\n ", g.CurrentTurn.Name, g.CurrentTurn.Id)
		steps := g.Dice.GenerateNumber()

		fmt.Printf("Player Dice Rolled Up Got Value%d\n", steps)
		g.CurrentTurn.UpdateCurrentPos(g.CurrentTurn.GetCurrPos() + steps)

		fmt.Printf("Player Updated Position %d\n ", g.CurrentTurn.GetCurrPos())

		present, finalPosition := g.Board.IsSankePresent(g.CurrentTurn.GetCurrPos())
		if present {
			g.CurrentTurn.UpdateCurrentPos(finalPosition)
			fmt.Printf("oops Snake Present %d\n ", g.CurrentTurn.GetCurrPos())
		}
		present, finalPosition = g.Board.IsLadderPresent(g.CurrentTurn.GetCurrPos())
		if present {
			g.CurrentTurn.UpdateCurrentPos(finalPosition)
			fmt.Printf("Ladder present moving to position %d\n ", g.CurrentTurn.GetCurrPos())
		}
		if g.CurrentTurn.GetCurrPos() >= 100 {
			fmt.Printf("Player won the game %s%d ", g.CurrentTurn.Name, g.CurrentTurn.Id)
			break
		}
		i++
		i = i % len
		g.CurrentTurn = g.Players[i]
	}
}

type Player struct {
	Id   int
	Name string
}

type PlayerPos struct {
	Player
	currPosition int
}

func NewPlayerPos(Player Player) *PlayerPos {
	return &PlayerPos{Player: Player, currPosition: 0}
}

func (p *PlayerPos) UpdateCurrentPos(val int) {
	p.currPosition = val
}

func (p *PlayerPos) GetCurrPos() int {
	return p.currPosition
}

func main() {
	Player1 := Player{
		Name: "Madhav",
		Id:   1,
	}
	Player2 := Player{
		Name: "Rahul",
		Id:   2,
	}
	d := NewDice(1)

	snakes := []Snake{{Start: 8, End: 5}}
	ladders := []Ladder{{Start: 12, End: 16}}

	Board := Board{
		Snakes:  snakes,
		Ladders: ladders,
	}
	g := NewGame(Board, []Player{Player1, Player2}, d)
	g.PlayGame()
}


```