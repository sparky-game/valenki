package parchis

import (
	valenkiCommon "github.com/sparky-game/valenki/pkg/common"
)

type Game struct {
	Players     []*Player
	Board       *Board
	CurrentTurn int
	DiceValue   int
	IsFinished  bool
}

func NewGame(playerNames []string) (*Game, error) {
	if len(playerNames) < 2 || len(playerNames) > 4 {
		return nil, errors.New("Invalid number of players, must be between 2 and 4")
	}
	players := make([]*Player, len(playerNames))
	colors := []string{"Red", "Blue", "Yellow", "Green"}
	for i, name := range playerNames {
		players[i] = NewPlayer(name, colors[i])
	}
	return &Game{
		Players:     players,
		Board:       NewBoard(),
		CurrentTurn: 0,
		IsFinished:  false,
	}, nil
}

func (g *Game) RollDice() int {
	g.DiceValue = valenkiCommon.RollDice(1)
	return g.DiceValue
}

func (g *Game) NextTurn() {
	g.CurrentTurn = (g.CurrentTurn + 1) % len(g.Players)
}

func (g *Game) GetCurrentPlayer() *Player {
	return g.Players[g.CurrentTurn]
}

func (g *Game) CheckIfGameFinished() bool {
	for _, player := range g.Players {
		if player.HasWon() {
			return true
		}
	}
	return false
}

func (g *Game) MovePiece(playerIdx int, pieceID int, n int) error {
	player := g.Players[playerIdx]
	if piece, err := player.GetPieceByID(pieceID); err != nil {
		return err
	}
	if piece.Position+n > 68 {
		return errors.New("You need the exact number to enter the goal")
	}
	g.Board.MovePiece(piece, n)
	if g.Board.CanCapture(piece, piece.Position) {
		if err := g.Board.CapturePiece(piece, piece.Position); err != nil {
			return err
		}
	}
	if player.HasWon() {
		g.IsFinished = true
	}
	return nil
}
