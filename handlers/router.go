package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `<h1>Hello, World!</h1>`)
}

func Router() http.Handler {
	router := chi.NewRouter()

	pages := PageController()
	router.Get("/", pages.Index)

	return router
}
