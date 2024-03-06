package main

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	db *mongo.Collection
}

const (
	port = ":8082"
	uri  = "mongodb://mongodb:27017"
)

func main() {

	app := Config{}
	app.ConnectDB()

	server := http.Server{
		Addr:    port,
		Handler: app.routes(),
	}

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
	app.db = client.Database("todo").Collection("user")

	fmt.Println("Connected to MongoDB")

}
