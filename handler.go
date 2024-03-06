package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var responseData map[string]interface{}

type ReadJson struct {
	Action  string         `json:"action"`
	Payload map[string]any `json:"payload"`
}

func (app *Config) handlerSubmission(w http.ResponseWriter, r *http.Request) {

	fmt.Println("installing config")
	var readJson ReadJson

	if err := json.NewDecoder(r.Body).Decode(&readJson); err != nil {
		fmt.Println("Error decoding")
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	fmt.Println("config", readJson)
	switch readJson.Action {
	case "register":
		RegisterUser(w, readJson)
	case "login":
		LoginUser(w, r, readJson)
	case "getme":
		app.GetProfile(w, r)
	case "createTask":
		app.CreateTask(w, r, readJson.Payload)

	case "allTask":
		app.GetAllTask(w, r)
	case "getTask":
		app.GetTask(w, r, readJson.Payload)
	case "updateTask":
		app.UpdateTask(w, r, readJson.Payload)

	case "deleteTask":
		app.DeleteTask(w, r, readJson.Payload)

	default:
		ErrorHandler(w, http.StatusBadRequest)
	}

}

func RegisterUser(w http.ResponseWriter, data ReadJson) {
	fmt.Println("entering register")
	jsonData, err := json.MarshalIndent(data.Payload, "", "\t")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	requset, err := http.NewRequest("POST", "http://user:8081/register", bytes.NewBuffer(jsonData))
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	fmt.Println("Register successfully")
	client := &http.Client{}

	res, err := client.Do(requset)
	fmt.Println("response.........", res.Body)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(res.Body).Decode(&responseData)
	fmt.Println("responsedatatatat........", responseData)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(responseData)

}

func LoginUser(w http.ResponseWriter, r *http.Request, data ReadJson) {

	jsonData, err := json.MarshalIndent(data.Payload, "", "\t")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	request, err := http.NewRequest("POST", "http://user:8081/login", bytes.NewBuffer(jsonData))
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil || res.StatusCode != 200 {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(res.Body).Decode(&responseData)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(responseData)
	defer r.Body.Close()
}

func (app *Config) GetProfile(w http.ResponseWriter, r *http.Request) {

	id, err := app.auth(r)
	if err != nil {
		ErrorHandler(w, http.StatusUnauthorized)
		return
	}

	fmt.Println("id: ", id)
	baseUrl := "http://user:8081/me"
	url := fmt.Sprintf("%s/%s", baseUrl, id)
	fmt.Println("Request url", url)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(res.Body).Decode(&responseData)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	fmt.Println("response", responseData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
	defer r.Body.Close()
}

func (app *Config) auth(r *http.Request) (string, error) {
	request, err := http.NewRequest("GET", "http://auth:8082/", nil)
	if err != nil {

		return "", err
	}
	for key, values := range r.Header {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {

		return "", err
	}
	if res.StatusCode == 401 || res.StatusCode == 500 {
		fmt.Println("Invalid")
		return "", err
	}
	err = json.NewDecoder(res.Body).Decode(&responseData)
	if err != nil {
		return "", err
	}
	fmt.Println("response", responseData)
	return responseData["id"].(string), nil

}

// tasks
