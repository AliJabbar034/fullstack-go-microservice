package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", app.CreateTaskHandler)
	r.Get("/{userId}", app.GetAllTaskHandler)
	r.Get("/task/{taskId}", app.GetTaskHandler)
	r.Put("/task/{taskId}/{userId}", app.UpdateTaskHandler)
	r.Delete("/task/{taskId}/{userId}", app.DeleteTaskHandler)

	return r
}
