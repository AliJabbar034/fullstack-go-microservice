package main

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	port = ":8083"
	uri  = "mongodb://mongodb:27017"
)

type Config struct {
	db *mongo.Collection
}

func main() {
	app := Config{}

	app.ConnectDB()
	server := http.Server{
		Addr:    port,
		Handler: app.routes(),
	}
	fmt.Println("Starting task manager...........")
	server.ListenAndServe()
}

func (app *Config) ConnectDB() {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// cancel()
	// if err := client.Ping(ctx, nil); err != nil {
	// 	panic(err.Error())
	// }
	app.db = client.Database("todo").Collection("tasks")

	fmt.Println("Connected to MongoDB")

}
