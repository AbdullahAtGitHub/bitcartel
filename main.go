package main

import (
	"log"
	"net/http"

	"github.com/AbdullahAtGitHub/bitcartal/handlers"
)

const PORT = ":8080"

func main() {
	router := handlers.Router()

	log.Printf("Starting server on %s", PORT)

	log.Fatal(http.ListenAndServe(PORT, router))
}
