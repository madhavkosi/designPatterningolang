**Solution:**
```go

package main

import "fmt"

type Player struct {
	PlayerId int
	Name     string
	Contact  string
}

func (p Player) DisplayInfo() {
	fmt.Printf("Player %d (%s)'s turn\n", p.PlayerId, p.Name)
}

func (p Player) PlayTurn() (int, int) {
	var i, j int
	for {
		fmt.Println("Enter position i (row):")
		fmt.Scanf("%d", &i)
		fmt.Println("Enter position j (col):")
		fmt.Scanf("%d", &j)
		if i >= 0 && i < 3 && j >= 0 && j < 3 {
			break
		}
		fmt.Println("Invalid input! Please enter values between 0 and 2.")
	}
	return i, j
}

type Symbols string

const (
	X Symbols = "X"
	O Symbols = "O"
)

type Game struct {
	Board         [][]Symbols
	Player1       Player
	Player2       Player
	CurrentPlayer *Player
	Symbol1       Symbols
	Symbol2       Symbols
}
type GameBuilder struct {
	G *Game
}

func NewGame(player1, player2 Player) *Game {
	board := make([][]Symbols, 3)
	for i := range board {
		board[i] = make([]Symbols, 3)
	}
	return &Game{
		Board:         board,
		Player1:       player1,
		Player2:       player2,
		CurrentPlayer: &player1,
		Symbol1:       X,
		Symbol2:       O,
	}
}

func (g *Game) PlayGame() string {
	for i := 0; i < 9; i++ {
		g.Display()
		g.CurrentPlayer.DisplayInfo()
		symbol := g.getCurrentSymbol()
		posI, posJ := g.GetValidPosition()
		g.Board[posI][posJ] = symbol
		if g.CheckIfWon(posI, posJ, symbol) {
			g.Display()
			return fmt.Sprintf("Player %d (%s) won!", g.CurrentPlayer.PlayerId, g.CurrentPlayer.Name)
		}
		g.SwitchPlayer()
	}
	g.Display()
	return "The game is a draw."
}

func (g *Game) getCurrentSymbol() Symbols {
	if g.CurrentPlayer.PlayerId == g.Player1.PlayerId {
		return g.Symbol1
	}
	return g.Symbol2
}

func (g *Game) GetValidPosition() (int, int) {
	var posI, posJ int
	for {
		posI, posJ = g.CurrentPlayer.PlayTurn()
		if g.IsValid(posI, posJ) {
			break
		}
		fmt.Println("Position is already taken, please choose another.")
	}
	return posI, posJ
}

func (g *Game) SwitchPlayer() {
	if g.CurrentPlayer.PlayerId == g.Player1.PlayerId {
		g.CurrentPlayer = &g.Player2
	} else {
		g.CurrentPlayer = &g.Player1
	}
}

func (g *Game) Display() {
	fmt.Println("Current board:")
	for _, row := range g.Board {
		for _, cell := range row {
			if cell == "" {
				fmt.Print("_ ")
			} else {
				fmt.Printf("%s ", cell)
			}
		}
		fmt.Println()
	}
}

func (g *Game) IsValid(i, j int) bool {
	return g.Board[i][j] != X && g.Board[i][j] != O
}

func (g *Game) CheckIfWon(posI, posJ int, symbol Symbols) bool {
	// Check row
	if g.Board[posI][0] == symbol && g.Board[posI][1] == symbol && g.Board[posI][2] == symbol {
		return true
	}
	// Check column
	if g.Board[0][posJ] == symbol && g.Board[1][posJ] == symbol && g.Board[2][posJ] == symbol {
		return true
	}
	// Check diagonals
	if g.Board[0][0] == symbol && g.Board[1][1] == symbol && g.Board[2][2] == symbol {
		return true
	}
	if g.Board[0][2] == symbol && g.Board[1][1] == symbol && g.Board[2][0] == symbol {
		return true
	}
	return false
}

func main() {
	player1 := Player{PlayerId: 1, Name: "Madhav"}
	player2 := Player{PlayerId: 2, Name: "Rahul"}
	game := NewGame(player1, player2)
	result := game.PlayGame()
	fmt.Println(result)
}

```