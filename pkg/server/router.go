package server

import (
  "net/http"
)

func APIRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloWorldHandler)
	mux.HandleFunc("/roll", RollHandler)
	return JSONMiddleware(mux)
}
