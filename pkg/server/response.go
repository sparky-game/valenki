package server

import (
	valenkiCommon "github.com/sparky-game/valenki/pkg/common"
	"net/http"
	"time"
)

type Response struct {
	Timestamp      time.Time `json:"timestamp"`
	Message        string    `json:"message"`
	DiceRollResult int       `json:"dice_roll_result"`
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Timestamp: time.Now(),
		Message:   "Привет, Валенки!",
	}
	JSONEncoder(w, res)
}

func RollHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Timestamp:      time.Now(),
		Message:        "Rolled dice successfully!",
		DiceRollResult: valenkiCommon.RollDice(1),
	}
	JSONEncoder(w, res)
}
