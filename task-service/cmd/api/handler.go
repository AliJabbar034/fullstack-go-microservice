package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseJson struct {
	Data       any `json:"data"`
	StatusCode int `json:"status_code"`
}

func (app *Config) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println("error bod", err.Error())
		ErrorHandler(w, http.StatusBadRequest, errors.New("bad request: "))
		return
	}
	newTask, err := NewTask(task)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest, err)
		return
	}
	taskId, err := app.CreateTask(*newTask)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	response := ResponseJson{
		Data:       taskId,
		StatusCode: http.StatusCreated,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)
}

func (app *Config) GetAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "userId")

	fmt.Println("GetAllTaskHandler", id)

	tasks, err := app.GetAllTask(id)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	response := ResponseJson{
		Data:       tasks,
		StatusCode: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}

func (app *Config) GetTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "taskId")
	taskId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest, errors.New("invalid task id"))
		return
	}

	task, err := app.GetTask(taskId)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	taskResponse := &ResponseJson{
		Data:       task,
		StatusCode: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&taskResponse)

}

func (app *Config) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "taskId")
	userId := chi.URLParam(r, "userId")
	taskId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest, errors.New("invalid task id"))
		return
	}

	deleted, err := app.DeleteTask(userId, taskId)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	response := ResponseJson{
		Data:       deleted,
		StatusCode: http.StatusNoContent,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}

func (app *Config) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "taskId")
	userId := chi.URLParam(r, "userId")
	taskId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest, errors.New("invalid task id"))
		return
	}

	updated, err := app.UpdateTask(taskId, userId)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	response := &ResponseJson{
		Data:       updated,
		StatusCode: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&response)
}
