package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Message:   "Привет, Валенки!",
		Timestamp: time.Now(),
	}
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloWorldHandler)
	return JSONMiddleware(mux)
}

func main() {
	log.Println("Starting Валенки server on :6969 ...")
	if err := http.ListenAndServe(":6969", NewRouter()); err != nil {
		log.Fatalf("Unable to start server (%v)\n", err)
	}
}
