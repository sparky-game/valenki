package parchis

import (
	"errors"
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
	if g.DiceValue == 0 {
		g.DiceValue = valenkiCommon.RollDice(1)
	}
	return g.DiceValue
}

func (g *Game) NextTurn() {
	g.DiceValue = 0
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

func (g *Game) MovePiece(playerIdx int, pieceID int) error {
	if playerIdx < 0 || playerIdx >= len(g.Players) {
		return errors.New("Player ID doesn't exist")
	}
	if playerIdx != g.CurrentTurn {
		return errors.New("It's not your turn!")
	}
	player := g.Players[playerIdx]
	piece, err := player.GetPieceByID(pieceID)
	if err != nil {
		return err
	}
	if g.DiceValue == 0 {
		return errors.New("How do I move? Need to roll dice first!")
	}
	if piece.Position+g.DiceValue > 68 {
		return errors.New("You need the exact number to enter the goal")
	}
	g.Board.MovePiece(piece, g.DiceValue)
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
