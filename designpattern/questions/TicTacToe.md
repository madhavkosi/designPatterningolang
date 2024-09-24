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



For a Tic-Tac-Toe game like this, we can utilize design patterns to structure the code in a more scalable and maintainable way. Here are a few design patterns that can be applied:

### 1. **Strategy Pattern**:
   The `Strategy Pattern` can be used to separate the logic of how a player takes their turn. This is useful if we want to extend the game in the future by adding different types of players (e.g., human players, AI players with different difficulty levels).

   - **When to Use**: When we expect to have different strategies for how a player makes a move.
   - **How**: Create an interface `MoveStrategy` that defines how to get the next move, and different implementations for human players, AI players, etc.

   ```go
   type MoveStrategy interface {
       GetMove(board [][]Symbols) (int, int)
   }

   type HumanMove struct{}

   func (hm *HumanMove) GetMove(board [][]Symbols) (int, int) {
       var i, j int
       fmt.Println("Enter position i (row):")
       fmt.Scanf("%d", &i)
       fmt.Println("Enter position j (col):")
       fmt.Scanf("%d", &j)
       return i, j
   }

   type AIMove struct{}

   func (am *AIMove) GetMove(board [][]Symbols) (int, int) {
       // AI logic here for picking the move
       return 1, 1
   }
   ```

   - **Integration**: Each `Player` can have a `MoveStrategy`, and the player's `PlayTurn()` can delegate the logic to that strategy.

   ```go
   type Player struct {
       PlayerId      int
       Name          string
       MoveStrategy  MoveStrategy
   }

   func (p *Player) PlayTurn(board [][]Symbols) (int, int) {
       return p.MoveStrategy.GetMove(board)
   }
   ```

   This allows easy switching between different types of players, whether they are human or AI.

### 2. **Observer Pattern**:
   The `Observer Pattern` can be used to notify observers about the game's state. For example, if we want to add logging, score tracking, or display updates when a move is made, these can be done using observers.

   - **When to Use**: When we want to notify multiple components (UI, logging, etc.) about changes in the game state.
   - **How**: Define a `GameObserver` interface with a method `Update`, and have concrete observers like `GameUI`, `GameLogger` that implement the interface.

   ```go
   type GameObserver interface {
       Update(board [][]Symbols)
   }

   type GameUI struct{}

   func (ui *GameUI) Update(board [][]Symbols) {
       // Code to update UI
   }

   type GameLogger struct{}

   func (logger *GameLogger) Update(board [][]Symbols) {
       // Code to log the game state
   }
   ```

   - **Integration**: In the `Game` class, maintain a list of observers, and call their `Update` method whenever the game state changes.

   ```go
   type Game struct {
       Board      [][]Symbols
       Observers  []GameObserver
   }

   func (g *Game) AddObserver(o GameObserver) {
       g.Observers = append(g.Observers, o)
   }

   func (g *Game) NotifyObservers() {
       for _, o := range g.Observers {
           o.Update(g.Board)
       }
   }

   func (g *Game) MakeMove(i, j int, symbol Symbols) {
       g.Board[i][j] = symbol
       g.NotifyObservers()
   }
   ```

   This allows for decoupling the game logic from the UI and other concerns like logging.

### 3. **Factory Pattern**:
   The `Factory Pattern` can be used to create different types of players depending on user input or game mode.

   - **When to Use**: When you want to dynamically create players based on input (e.g., whether a player is a human or an AI).
   - **How**: Create a `PlayerFactory` that returns a player object based on some criteria (such as player type).

   ```go
   type PlayerType int

   const (
       Human PlayerType = iota
       AI
   )

   type PlayerFactory struct{}

   func (pf *PlayerFactory) CreatePlayer(id int, name string, playerType PlayerType) Player {
       switch playerType {
       case Human:
           return Player{PlayerId: id, Name: name, MoveStrategy: &HumanMove{}}
       case AI:
           return Player{PlayerId: id, Name: name, MoveStrategy: &AIMove{}}
       }
       return Player{}
   }
   ```

   - **Integration**: Use the factory to create players when setting up the game.

   ```go
   playerFactory := PlayerFactory{}
   player1 := playerFactory.CreatePlayer(1, "Madhav", Human)
   player2 := playerFactory.CreatePlayer(2, "AI Bot", AI)
   ```

   This makes adding new player types (such as different AI difficulties) easy without modifying existing code.

### 4. **State Pattern**:
   The `State Pattern` can be used to represent different game states (e.g., "Playing", "GameOver", "Draw"). Each state can encapsulate its behavior and what transitions are allowed.

   - **When to Use**: When you want to manage different game states cleanly, especially in larger games where each state may have specific behavior.
   - **How**: Define a `GameState` interface with methods like `HandleInput`, `Update`, etc., and create concrete states for "Playing", "GameOver", and "Draw".

   ```go
   type GameState interface {
       HandleInput(g *Game)
       Update(g *Game)
   }

   type PlayingState struct{}

   func (ps *PlayingState) HandleInput(g *Game) {
       // handle input
   }

   func (ps *PlayingState) Update(g *Game) {
       // check for game over
   }

   type GameOverState struct{}

   func (gs *GameOverState) HandleInput(g *Game) {
       fmt.Println("Game is over!")
   }

   func (gs *GameOverState) Update(g *Game) {
       // no updates allowed in game over
   }
   ```

   - **Integration**: The game holds a reference to the current state and delegates behavior to it.

   ```go
   type Game struct {
       State GameState
   }

   func (g *Game) SetState(state GameState) {
       g.State = state
   }

   func (g *Game) HandleInput() {
       g.State.HandleInput(g)
   }
   ```

### Summary of Design Patterns:
1. **Strategy Pattern**: To handle different strategies for player moves (human, AI).
2. **Observer Pattern**: To decouple game logic from display/logging/other external concerns.
3. **Factory Pattern**: To dynamically create players based on input (human or AI).
4. **State Pattern**: To manage different game states (playing, game over, draw).

Depending on how complex you want the Tic-Tac-Toe game to become, one or more of these patterns can be applied to make the code more scalable, maintainable, and extensible.