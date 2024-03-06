package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello  From User microservice ....."))
	})
	r.Post("/register", app.RegisterHandler)
	r.Post("/login", app.LoginUserHandler)
	r.Get("/me/{id}", app.GetProfile)

	return r
}
