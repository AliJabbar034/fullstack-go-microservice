package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var taskData map[string]interface{}

func (app *Config) CreateTask(w http.ResponseWriter, r *http.Request, data map[string]any) {

	id, err := app.auth(r)
	if err != nil {
		ErrorHandler(w, http.StatusUnauthorized)
		return
	}
	data["user_id"] = id

	fmt.Println("req data", data)

	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	req, err := http.NewRequest("POST", "http://task:8083/", bytes.NewBuffer(jsonData))
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(res.Body).Decode(&taskData)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&taskData)
	r.Body.Close()

}

func (app *Config) GetAllTask(w http.ResponseWriter, r *http.Request) {
	id, err := app.auth(r)
	if err != nil {
		ErrorHandler(w, http.StatusUnauthorized)
		return
	}
	baseUrl := "http://task:8083"
	url := fmt.Sprintf("%s/%s", baseUrl, id)
	fmt.Println("Request", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(res.Body).Decode(&taskData)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&taskData)

}

func (app *Config) GetTask(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {

	taskId := data["task_id"]
	baseUrl := "http://task:8083/task"
	url := fmt.Sprintf("%s/%s", baseUrl, taskId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(res.Body).Decode(&taskData)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&taskData)
	defer r.Body.Close()
}

func (app *Config) UpdateTask(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {

	userId, err := app.auth(r)
	if err != nil {
		ErrorHandler(w, http.StatusUnauthorized)
		return
	}
	taskId := data["task_id"]
	baseUrl := "http://task:8083/task"
	url := fmt.Sprintf("%s/%s/%s", baseUrl, taskId, userId)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(res.Body).Decode(&taskData)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&taskData)
	r.Body.Close()
}

func (app *Config) DeleteTask(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {

	userId, err := app.auth(r)
	if err != nil {
		ErrorHandler(w, http.StatusUnauthorized)
		return
	}
	taskId := data["task_id"]
	baseUrl := "http://task:8083/task"
	url := fmt.Sprintf("%s/%s/%s", baseUrl, taskId, userId)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(res.Body).Decode(&taskData)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&taskData)
	r.Body.Close()
}
