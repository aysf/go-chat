package main

import (
	"log"
	"net/http"

	"github.com/aysf/go-chat/internal/handlers"
)

func main() {

	mux := routes()

	log.Println("starting web server on port 4000")
	go handlers.ListenToWsChannel()

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Println(err)
	}
}
