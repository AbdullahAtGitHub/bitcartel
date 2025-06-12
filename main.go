package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const PORT = ":8080"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<h1>Hello, World!</h1>`)
}

func main() {
	router := chi.NewRouter()

	router.Get("/", handleIndex)

	log.Printf("Starting server on :%s", PORT)

	log.Fatal(http.ListenAndServe(PORT, router))
}
