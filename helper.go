package main

import (
	"encoding/json"
	"net/http"
)

type ErrorJson struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func ErrorHandler(w http.ResponseWriter, status int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorJson := &ErrorJson{
		Error: http.StatusText(status),
		Code:  status,
	}
	json.NewEncoder(w).Encode(&errorJson)
	return
}
