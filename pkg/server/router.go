package server

import (
	gmux "github.com/gorilla/mux"
	"net/http"
)

func APIRouter() http.Handler {
	mux := gmux.NewRouter()
	mux.HandleFunc("/", HelloWorldHandler).Methods("GET") // TODO: remove it
	mux.HandleFunc("/roll", RollHandler).Methods("GET")   // TODO: remove it
	// Parchis
	mux.HandleFunc("/parchis", ParchisHandlerCreateGame).Methods("POST")
	mux.HandleFunc("/parchis/{game_id}", ParchisHandlerGetGameState).Methods("GET")
	mux.HandleFunc("/parchis/{game_id}/roll", ParchisHandlerRollDice).Methods("POST")
	mux.HandleFunc("/parchis/{game_id}/move", ParchisHandlerMovePiece).Methods("POST")
	return JSONMiddleware(mux)
}
