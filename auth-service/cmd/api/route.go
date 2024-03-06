package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", app.AuthenticateUser)
	return r
}
