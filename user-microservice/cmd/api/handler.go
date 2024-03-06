package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseJson struct {
	Data  any    `json:"data"`
	Token string `json:"token"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var user User
	fmt.Println("Registering handler for config file ")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("error bod", err.Error())
		ErrorHandler(w, http.StatusBadRequest, errors.New("bad request: "))
		return
	}

	newUser, err := NewUser(user)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println("password hasing", newUser.Password)
	pass, err := HasPassword(newUser.Password)
	fmt.Println("password", pass)
	if err != nil {
		log.Println("error bod", err.Error())
		ErrorHandler(w, http.StatusBadRequest, errors.New("bad request: "))
		return
	}
	newUser.Password = pass
	fmt.Println("newUser", newUser)

	insertedId, err := app.db.InsertOne(context.Background(), &newUser)
	fmt.Println("insertedId", insertedId)

	if err != nil {
		log.Println("error insert", err.Error())
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	id := insertedId.InsertedID.(primitive.ObjectID).Hex()
	token, err := GenerateToken(id)
	if err != nil {
		log.Println("error token", err.Error())
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	responseJson := &ResponseJson{
		Data:  id,
		Token: token,
	}
	fmt.Println("u..................updated", responseJson)

	json.NewEncoder(w).Encode(&responseJson)

}

func (app *Config) LoginUserHandler(w http.ResponseWriter, r *http.Request) {

	var loginData LoginData

	fmt.Println("Login handler for config file ")

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		log.Println("error bod", err.Error())
		ErrorHandler(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	fmt.Println("loginData", loginData)

	var user User
	if err := app.db.FindOne(context.Background(), bson.M{"email": loginData.Email}).Decode(&user); err != nil {
		log.Println("error find", err.Error())
		ErrorHandler(w, http.StatusInternalServerError, errors.New("invalid login data"))
		return
	}

	fmt.Println("Comaparrerr", user)
	err := ComapreHasPass(user.Password, loginData.Password)
	fmt.Println("compare done")
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest, errors.New("invalid login data"))
		return
	}

	id := user.ID.Hex()
	token, err := GenerateToken(id)
	if err != nil {
		log.Println("error token", err.Error())
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	responseData := &ResponseJson{
		Data:  user,
		Token: token,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&responseData)

}

func (app *Config) GetProfile(w http.ResponseWriter, r *http.Request) {
	var user User
	id := chi.URLParam(r, "id")
	_id, _ := primitive.ObjectIDFromHex(id)

	if err := app.db.FindOne(context.Background(), bson.M{"_id": _id}).Decode(&user); err != nil {
		log.Println("error find", err.Error())
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	responseData := &ResponseJson{
		Data: user,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&responseData)
}
