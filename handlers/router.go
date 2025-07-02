package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	fs := http.Dir("./assets/dist")
	router.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(fs)))

	pages := PageController()
	router.Get("/", pages.Index)

	return router
}
