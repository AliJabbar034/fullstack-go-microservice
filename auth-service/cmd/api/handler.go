package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}
type Response struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

func (app *Config) AuthenticateUser(w http.ResponseWriter, r *http.Request) {

	tokenstring := r.Header.Get("Authorization")
	if tokenstring == "" {
		ErrorHandler(w, http.StatusUnauthorized, errors.New("unauthorized"))

		return
	}
	parts := strings.Split(tokenstring, " ")
	if len(parts) != 2 {
		ErrorHandler(w, http.StatusBadRequest, errors.New("invalid token"))
		return
	}
	token := parts[1]
	fmt.Println("token", token)
	if token == "" {
		ErrorHandler(w, http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	tokenstr, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte("raggytwwir5w5ww5"), nil
	})

	fmt.Println("GGGGGGGG", tokenstr)
	if err != nil {
		fmt.Println(err.Error())
		ErrorHandler(w, http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	claims, ok := tokenstr.Claims.(jwt.MapClaims)
	fmt.Println("claims: ", claims)
	if !ok && !tokenstr.Valid {
		ErrorHandler(w, http.StatusUnauthorized, errors.New("invalid token"))
		return
	}
	claimId, ok := claims["id"]
	if !ok {
		ErrorHandler(w, http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	fmt.Println("claimId: ", claimId)
	id := claimId.(string)
	fmt.Println("token.......", id)
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest, errors.New("invalid token"))
		return
	}
	var user User
	err = app.db.FindOne(context.Background(), bson.M{"_id": _id}).Decode(&user)
	if err != nil {

		fmt.Println("Error decoding", err.Error())
		ErrorHandler(w, http.StatusUnauthorized, errors.New("invalid token"))
		return
	}
	fmt.Println("token  returned ")

	response := &Response{
		Message: "success",
		ID:      id,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}
