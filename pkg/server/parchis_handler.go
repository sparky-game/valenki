package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sparky-game/valenki/pkg/parchis"
	"net/http"
	"strconv"
)

// In-memory store of Parchis games (for simplicity, no database)
var games = make(map[string]*parchis.Game)

func ParchisHandlerCreateGame(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Players []string `json:"players"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	game, err := parchis.NewGame(data.Players)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gameID := strconv.Itoa(len(games) + 1)
	games[gameID] = game
	JSONEncoder(w, http.StatusOK, map[string]string{"game_id": gameID})
}

func ParchisHandlerGetGameState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game, ok := games[vars["game_id"]]
	if !ok {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}
	res := struct {
		Players     []*parchis.Player `json:"players"`
		CurrentTurn int               `json:"current_turn"`
		IsFinished  bool              `json:"is_finished"`
	}{
		Players:     game.Players,
		CurrentTurn: game.CurrentTurn,
		IsFinished:  game.IsFinished,
	}
	JSONEncoder(w, http.StatusOK, res)
}

func ParchisHandlerRollDice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game, ok := games[vars["game_id"]]
	if !ok {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}
	JSONEncoder(w, http.StatusOK, map[string]int{"dice_roll": game.RollDice()})
}

func ParchisHandlerMovePiece(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game, ok := games[vars["game_id"]]
	if !ok {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}
	var data struct {
		PlayerID int `json:"player_id"`
		PieceID  int `json:"piece_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := game.MovePiece(data.PlayerID, data.PieceID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	game.NextTurn()
	JSONEncoder(w, http.StatusOK, map[string]string{"status": "moved successfully"})
}
