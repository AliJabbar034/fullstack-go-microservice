package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":8000"
)

type Config struct {
}

func main() {

	app := Config{}
	serve := http.Server{
		Addr:    port,
		Handler: app.routes(),
	}
	fmt.Println("Starting.........")
	serve.ListenAndServe()
}
