package main

import (
	"log"
	"net/http"
)

func main() {

	mux := routes()

	log.Println("starting web server on port 4000")

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Println(err)
	}
}
