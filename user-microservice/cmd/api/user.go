package main

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email,"`
	Password string             `bson:"password" json:"password"`
}

func NewUser(user User) (*User, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return nil, errors.New("bad request")
	}
	return &User{
		ID:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
