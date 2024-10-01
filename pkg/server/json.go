package server

import (
	"encoding/json"
	"net/http"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func JSONEncoder(w http.ResponseWriter, status int, r interface{}) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	w.WriteHeader(status)
	if err := encoder.Encode(r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
