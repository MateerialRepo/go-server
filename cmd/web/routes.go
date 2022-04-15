package main

import (
	"github.com/MateerialRepo/go-webapp/pkg/config"
	"github.com/MateerialRepo/go-webapp/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	//Routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/", handlers.Repo.About)

	return mux
}
