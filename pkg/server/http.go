package server

import (
  "fmt"
  "net/http"
  "log"
)

func HTTP(port int) {
	log.Printf("Starting the Валенки server on :%d ...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), APIRouter()); err != nil {
		log.Fatalf("Unable to start the server (%v)\n", err)
	}
}
