package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"io/ioutil"
	"net/http"
)

type CRUDHandler struct {
	cookieStore *sessions.CookieStore
	storage     StorageInterface
}

type GetPasswordRequest struct {
	PasswordIdentifier string `json:"id"`
}

type UserRequest struct {
	Name string `json:"name"`
}

type PasswordRequest struct {
	Password string `json:"password"`
	Id       string `json:"id"`
}

func (handler CRUDHandler) GetPassword(writer http.ResponseWriter, request *http.Request) {
	user := handler.storage.GetUser(request.Form.Get("UserId"))
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer request.Body.Close()
	var passwordId GetPasswordRequest
	err = json.Unmarshal(b, &passwordId)
	password, err := handler.storage.GetPassword(user, passwordId.PasswordIdentifier)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	/*
		Send the password "plain" as received from the database, Encryption and Decryption in frontend
	*/
	passwordJson, err := json.Marshal(password)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprint(writer, string(passwordJson))
}

func (handler CRUDHandler) AddPassword(writer http.ResponseWriter, request *http.Request) {
	user := handler.storage.GetUser(request.Form.Get("UserId"))
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer request.Body.Close()
	var password PasswordRequest
	err = json.Unmarshal(b, &password)
	err = handler.storage.AddPassword(user, password.Id, &Password{
		Password: password.Password,
		Id:       password.Id,
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	sendCRUDAnswer("CREATED", "", writer)
}

func (handler CRUDHandler) RemovePassword(writer http.ResponseWriter, request *http.Request) {
	user := handler.storage.GetUser(request.Form.Get("UserId"))
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer request.Body.Close()
	var passwordId GetPasswordRequest
	err = json.Unmarshal(b, &passwordId)
	err = handler.storage.DeletePassword(user, passwordId.PasswordIdentifier)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	sendCRUDAnswer("REMOVED", "", writer)
}

func (handler CRUDHandler) RemoveUser(writer http.ResponseWriter, request *http.Request) {
	user := handler.storage.GetUser(request.Form.Get("UserId"))
	err = handler.storage.RemoveUser(user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	sendCRUDAnswer("REMOVED", "", writer)
}

func (handler CRUDHandler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer request.Body.Close()
	var user UserRequest
	err = json.Unmarshal(b, &user)
	if user.Name == "" {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = handler.storage.UpdateUser(&User{
		Name: user.Name,
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	sendCRUDAnswer("UPDATED", "", writer)
}

func (handler CRUDHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	user := handler.storage.GetUser(request.Form.Get("UserId"))
	err = handler.storage.RemoveUser(user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	// Wrap struct around internal User struct to enforce string encoding of master password
	userObject := struct {
		Name           string
		MasterPassword string
	}{
		user.Name,
		string(user.MasterPassword),
	}
	userJson, err := json.Marshal(userObject)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(writer, string(userJson))
}

func sendCRUDAnswer(statusMessage string, errMessage string, writer http.ResponseWriter) {
	responseMessagJSON, err := json.Marshal(struct {
		Status string
		Error  string
	}{
		Status: statusMessage,
		Error:  errMessage,
	})
	checkError(err, writer)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(writer, string(responseMessagJSON))
}
