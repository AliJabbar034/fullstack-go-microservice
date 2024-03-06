package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type ErrorJson struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func ErrorHandler(w http.ResponseWriter, status int, err error) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorJson := &ErrorJson{
		Error: err.Error(),
		Code:  status,
	}
	json.NewEncoder(w).Encode(&errorJson)

}

func GenerateToken(id string) (string, error) {

	tonenstr := jwt.New(jwt.SigningMethodHS256)
	claims := tonenstr.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	token, err := tonenstr.SignedString([]byte("raggytwwir5w5ww5"))
	if err != nil {
		return "", err
	}
	return token, nil

}

func HasPassword(pass string) (string, error) {

	fmt.Println("if password", pass)
	passs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	fmt.Println("password done")
	return string(passs), nil
}

func ComapreHasPass(userPass string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(password))

	if err != nil {
		return err
	}
	return nil
}
