package server

import (
	gmux "github.com/gorilla/mux"
	"net/http"
)

func APIRouter() http.Handler {
	mux := gmux.NewRouter()
	mux.HandleFunc("/", HelloWorldHandler).Methods("GET") // TODO: remove it
	mux.HandleFunc("/roll", RollHandler).Methods("GET")   // TODO: remove it
	return JSONMiddleware(mux)
}
