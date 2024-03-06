package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Completed   bool               `bson:"completed" json:"completed"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UserID      string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
}

func NewTask(task Task) (*Task, error) {
	if task.Title == "" || task.Description == "" || task.UserID == "" {
		return nil, errors.New("bad request")
	}

	return &Task{
		ID:          primitive.NewObjectID(),
		Title:       task.Title,
		Description: task.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UserID:      task.UserID,
	}, nil
}

func (app *Config) CreateTask(task Task) (string, error) {

	inserted, err := app.db.InsertOne(context.Background(), &task)
	if err != nil {
		return "", err
	}
	id := inserted.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (app *Config) GetAllTask(userId string) ([]Task, error) {

	filter := bson.M{"user_id": userId}
	var tasks []Task
	cursor, err := app.db.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}
	fmt.Println("Completed", tasks)
	return tasks, nil

}

func (app *Config) GetTask(id primitive.ObjectID) (*Task, error) {

	var task Task

	filter := bson.M{"_id": id}
	if err := app.db.FindOne(context.Background(), filter).Decode(&task); err != nil {
		return nil, err
	}

	return &task, nil
}

func (app *Config) DeleteTask(userId string, taskId primitive.ObjectID) (int64, error) {

	filter := bson.M{"_id": taskId, "user_id": userId}
	deleted, err := app.db.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}

	return deleted.DeletedCount, nil
}

func (app *Config) UpdateTask(taskId primitive.ObjectID, userId string) (int64, error) {

	filter := bson.M{"_id": taskId, "user_id": userId}
	update := bson.M{"$set": bson.M{"completed": true}}
	updated, err := app.db.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}
	if updated.ModifiedCount == 0 {

		return 0, errors.New("task not found")
	}
	return updated.ModifiedCount, nil
}
