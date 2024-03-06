package main

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ErrorHandler(w http.ResponseWriter, statusCode int, err error) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorJson := &Error{
		Message: err.Error(),
		Code:    statusCode,
	}
	json.NewEncoder(w).Encode(&errorJson)
}
