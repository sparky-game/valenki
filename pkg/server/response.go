package server

import (
	valenkiCommon "github.com/sparky-game/valenki/pkg/common"
	"net/http"
	"time"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	res := struct {
		Timestamp time.Time `json:"timestamp"`
		Message   string    `json:"message"`
	}{
		Timestamp: time.Now(),
		Message:   "Привет, Валенки!",
	}
	JSONEncoder(w, http.StatusOK, res)
}

func RollHandler(w http.ResponseWriter, r *http.Request) {
	res := struct {
		Timestamp time.Time `json:"timestamp"`
		DiceRoll  int       `json:"dice_roll"`
	}{
		Timestamp: time.Now(),
		DiceRoll:  valenkiCommon.RollDice(1),
	}
	JSONEncoder(w, http.StatusOK, res)
}
