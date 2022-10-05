package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sanathshetty444/go-bookings/pkg/config"
	"github.com/sanathshetty444/go-bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(AddSessionManager)
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/divide", http.HandlerFunc(handlers.Repo.Divide))

	return mux
}
