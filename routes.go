package main

import (
	"net/http"

	"github.com/beejan2986/Bookings/pkg/config"
	"github.com/beejan2986/Bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// add some midleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	// routing
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
